package federation

import (
	"context"
	"crypto/tls"
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
		return "", "", config.NewConfigError(err)
	}

	callback, err := newCallbackHandler(logger)
	if err != nil {
		return "", "", err
	}

	err = callback.ListenAndServe()
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
		if writer != nil {
			if _, err = fmt.Fprintln(writer, text); err != nil {
				return "", "", fmt.Errorf("fprintln error: %w", err)
			}
		} else {
			logger.InfoContext(ctx, text)
		}
	} else {
		logger.DebugContext(ctx, "open browser", slog.String("url", authURL.String()))
		text := fmt.Sprintf(
			"Switch to your browser to complete the authentication process. "+
				"If it didn't open automatically, use the following link: %s",
			authURL.String(),
		)
		if writer != nil {
			if _, err = fmt.Fprintln(writer, text); err != nil {
				return "", "", fmt.Errorf("fprintln error: %w", err)
			}
		} else {
			logger.InfoContext(ctx, text)
		}
		browserErr = openBrowser(authURL.String())
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

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return GetTokenResult{}, err
	}
	defer resp.Body.Close() //nolint:errcheck

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
	tokenURL, err := url.Parse(federationURL + tokenEP)
	if err != nil {
		return GetTokenResult{}, config.NewConfigError(err)
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

	logger.DebugContext(ctx, "auth token is received")

	return token, nil
}

func openBrowser(url string) <-chan error {
	ch := make(chan error, 1)

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		if IsWSL() {
			cmd = exec.Command("cmd.exe", "/c", "start", strings.ReplaceAll(url, "&", "^&"))
			home, err := config.UserHomeDir()
			if err != nil {
				ch <- err
				return ch
			}
			cmd.Dir = home
		} else {
			cmd = exec.Command("xdg-open", url)
		}
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		ch <- fmt.Errorf("unsupported platform")
		return ch
	}

	go func() {
		output, err := cmd.CombinedOutput()
		if err != nil {
			ch <- errors.Join(err, errors.New(string(output)))
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
