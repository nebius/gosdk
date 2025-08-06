package federation

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sync"
)

type callbackHandler struct {
	server   *http.Server
	listener net.Listener
	serveErr chan error
	done     chan struct{}
	lock     sync.Mutex
	code     string
	state    string
}

func (h *callbackHandler) setCode(code, state string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	select {
	case <-h.done:
		// already set and closed
		return
	default:
		if code != "" && state == h.state {
			h.code = code
		}
		close(h.done)
	}
}

func (h *callbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	h.setCode(code, state)
	if h.code != "" && state == h.state {
		_, _ = fmt.Fprintln(w, "Login is successful, you may close the browser tab and go to the console")
	} else {
		_, _ = fmt.Fprintln(w, "Login is not successful, you may close the browser tab and try again")
	}
}

func (h *callbackHandler) Shutdown(ctx context.Context) error {
	err := h.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func newCallbackHandler(logger *slog.Logger) (*callbackHandler, error) {
	state, err := GenerateRandomString()
	if err != nil {
		return nil, err
	}

	handler := &callbackHandler{
		server: &http.Server{ //nolint:gosec // it's fine to have no ReadHeaderTimeout on localhost
			Handler:  nil, // is set below
			ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
		},
		listener: nil, // will be set in [callbackHandler.Serve]
		serveErr: make(chan error, 1),
		done:     make(chan struct{}),
		lock:     sync.Mutex{},
		code:     "", // the result of callback
		state:    state,
	}
	handler.server.Handler = handler
	return handler, nil
}

func (h *callbackHandler) ListenAndServe(ctx context.Context) error {
	var err error
	if h.listener, err = listener(ctx); err != nil {
		return err
	}

	go func() {
		err = h.server.Serve(h.listener)
		if !errors.Is(err, http.ErrServerClosed) {
			h.serveErr <- err
		}
	}()

	return nil
}

func listener(ctx context.Context) (net.Listener, error) {
	var errs error
	var config net.ListenConfig
	l, err := config.Listen(ctx, "tcp", "127.0.0.1:0")
	if err != nil {
		errs = err
		if l, err = config.Listen(ctx, "tcp6", "[::1]:0"); err != nil {
			errs = errors.Join(errs, err)
		}
	}
	if errs != nil {
		return l, errs
	}
	return l, nil
}
