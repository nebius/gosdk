package federation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strings"

	"github.com/nebius/gosdk/config"
	"github.com/nebius/gosdk/config/paths"
)

func getCode(
	ctx context.Context,
	clientID, authEndpoint, federationID string,
	pkceCode PKCECode,
	logger *slog.Logger,
	writer io.Writer,
	noBrowserOpen bool,
) (_ string, _ string, retErr error) {
	authURL, err := url.Parse(httpsURL(authEndpoint))
	if err != nil {
		return "", "", config.NewError(err)
	}

	callback, err := newCallbackHandler(logger)
	if err != nil {
		return "", "", err
	}

	err = callback.ListenAndServe(ctx)
	if err != nil {
		return "", "", err
	}
	defer func() {
		errX := callback.Shutdown(ctx)
		if errX != nil {
			retErr = errors.Join(retErr, errX)
		}
	}()

	redirectURI := url.URL{Scheme: "http", Host: callback.listener.Addr().String()}

	query := authURL.Query()
	query.Set("response_type", "code")
	if federationID != "" {
		query.Set("federation-id", federationID)
	}
	query.Set("client_id", clientID)
	query.Set("scope", "openid")
	query.Set("state", callback.state)
	query.Set("redirect_uri", redirectURI.String())
	query.Set("code_challenge", pkceCode.Challenge())
	query.Set("code_challenge_method", pkceCode.Method())
	authURL.RawQuery = query.Encode()

	var browserErr <-chan error
	if noBrowserOpen {
		text := fmt.Sprintf(
			"To complete the authentication process, open the following link "+
				"in your browser: %s",
			authURL.String(),
		)
		logger.DebugContext(
			ctx,
			"no browser flag is found, asking user to open auth link",
			slog.String("url", authURL.String()),
		)
		err = writeOrLog(ctx, logger, writer, text)
		if err != nil {
			return "", "", err
		}
	} else {
		logger.DebugContext(ctx, "open browser", slog.String("url", authURL.String()))
		text := fmt.Sprintf(
			"Switch to your browser to complete the authentication process. "+
				"If it didn't open automatically, use the following link: %s",
			authURL.String(),
		)
		err = writeOrLog(ctx, logger, writer, text)
		if err != nil {
			return "", "", err
		}
		browserErr = openBrowser(ctx, logger, authURL.String())
	}

	select {
	case <-callback.done:
		return callback.code, redirectURI.String(), nil
	case err = <-browserErr:
		return "", "", fmt.Errorf("open browser: %w", err)
	case err = <-callback.serveErr:
		return "", "", fmt.Errorf("server callback handler: %w", err)
	case <-ctx.Done():
		return "", "", ctx.Err()
	}
}

func writeOrLog(ctx context.Context, logger *slog.Logger, writer io.Writer, text string) error {
	if writer == nil {
		logger.InfoContext(ctx, text)
		return nil
	}
	_, err := fmt.Fprintln(writer, text)
	if err != nil {
		return fmt.Errorf("fprintln error: %w", err)
	}
	return nil
}

type GetTokenResult struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func getToken(
	ctx context.Context,
	clientID, tokenURL, code, redirectURI, verifier string,
) (GetTokenResult, error) {
	query := url.Values{}
	query.Set("grant_type", "authorization_code")
	query.Set("code", code)
	query.Set("redirect_uri", redirectURI)
	query.Set("client_id", clientID)
	query.Set("code_verifier", verifier)

	req, err := http.NewRequestWithContext(
		ctx, http.MethodPost, tokenURL,
		strings.NewReader(query.Encode()),
	)
	if err != nil {
		return GetTokenResult{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return GetTokenResult{}, err
	}
	defer func() { _ = resp.Body.Close() }()

	var result GetTokenResult
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return GetTokenResult{}, err
	}
	return result, nil
}

func Authorize(
	ctx context.Context,
	clientID, federationURL, federationID string,
	logger *slog.Logger,
	writer io.Writer,
	noBrowserOpen bool,
) (GetTokenResult, error) {
	logger = logger.With(
		slog.String("name", "federation_authorize"),
	)
	tokenURL, err := url.Parse(federationURL + tokenEP)
	if err != nil {
		return GetTokenResult{}, config.NewError(err)
	}

	pkceCode, err := GeneratePKCECode()
	if err != nil {
		return GetTokenResult{}, err
	}

	authURL := federationURL + authEP
	code, redirectURI, err := getCode(
		ctx,
		clientID, authURL, federationID,
		pkceCode, logger, writer, noBrowserOpen,
	)
	if err != nil {
		return GetTokenResult{}, fmt.Errorf("get code: %w", err)
	}

	logger.DebugContext(ctx, "auth code is received")

	token, err := getToken(
		ctx,
		clientID,
		httpsURL(tokenURL.String()),
		code, redirectURI,
		pkceCode.Verifier(),
	)
	if err != nil {
		return GetTokenResult{}, fmt.Errorf("exchange code to token: %w", err)
	}

	return token, nil
}

func openBrowser(ctx context.Context, logger *slog.Logger, url string) <-chan error {
	ch := make(chan error, 1)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		if IsWSL() {
			//nolint:gosec // we trust url
			cmd = exec.CommandContext(ctx, "cmd.exe", "/c", "start", strings.ReplaceAll(url, "&", "^&"))
			home, err := paths.UserHomeDir()
			if err != nil {
				ch <- err
				return ch
			}
			cmd.Dir = home
		} else {
			providers := []string{"xdg-open", "x-www-browser", "www-browser"}
			for _, provider := range providers {
				_, err := exec.LookPath(provider)
				if err != nil {
					continue
				}
				cmd = exec.CommandContext(ctx, provider, url)
				break
			}
			if cmd == nil {
				logger.DebugContext(
					ctx,
					"browser provider not found, choosing xdg-open to fail",
				)
				cmd = exec.CommandContext(ctx, "xdg-open", url)
			}
		}
	case "freebsd", "openbsd", "netbsd":
		cmd = exec.CommandContext(ctx, "xdg-open", url)
	case "windows":
		cmd = exec.CommandContext(ctx, "rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.CommandContext(ctx, "open", url)
	default:
		ch <- errors.New("unsupported platform")
		return ch
	}
	logger.DebugContext(ctx, "exec browser command", slog.String("cmd", cmd.String()))

	go func() {
		output, err := cmd.CombinedOutput()
		if err != nil {
			logger.ErrorContext(ctx, "open browser command failed",
				slog.String("output", string(output)),
				slog.String("error", err.Error()),
			)
			ch <- errors.Join(err, errors.New(string(output)))
		} else {
			logger.DebugContext(ctx, "open browser command completed", slog.String("output", string(output)))
		}
	}()

	return ch
}

func httpsURL(url string) string {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "https://" + url
	}
	return url
}
