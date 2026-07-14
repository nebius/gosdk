package browser

import (
	"context"
	"errors"
	"log/slog"
	"os/exec"
	"runtime"
	"strings"

	"github.com/nebius/gosdk/config/paths"
	"github.com/nebius/gosdk/platform"
)

func Command(ctx context.Context, url string) (*exec.Cmd, error) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		if platform.IsWSL() {
			//nolint:gosec // URL is passed as a browser argument.
			cmd = exec.CommandContext(ctx, "cmd.exe", "/c", "start", strings.ReplaceAll(url, "&", "^&"))
			home, err := paths.UserHomeDir()
			if err != nil {
				return nil, err
			}
			cmd.Dir = home
			return cmd, nil
		}
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
			return nil, errors.New("browser provider not found")
		}
	case "freebsd", "openbsd", "netbsd":
		cmd = exec.CommandContext(ctx, "xdg-open", url)
	case "windows":
		cmd = exec.CommandContext(ctx, "rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.CommandContext(ctx, "open", url)
	default:
		return nil, errors.New("unsupported platform")
	}
	_, err := exec.LookPath(cmd.Path)
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func Run(ctx context.Context, logger *slog.Logger, cmd *exec.Cmd) error {
	if logger != nil {
		logger.DebugContext(ctx, "exec browser command", slog.String("cmd", cmd.String()))
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		if logger != nil {
			logger.ErrorContext(ctx, "open browser command failed",
				slog.String("output", string(output)),
				slog.String("error", err.Error()),
			)
		}
		return errors.Join(err, errors.New(string(output)))
	}
	if logger != nil {
		logger.DebugContext(ctx, "open browser command completed", slog.String("output", string(output)))
	}
	return nil
}
