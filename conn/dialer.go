package conn

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"golang.org/x/sync/singleflight"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var ErrDialerClosed = errors.New("dialer is closed")

type addressDialOption struct {
	grpc.EmptyDialOption
	address Address
	opts    []grpc.DialOption
}

// WithAddressDialOptions allows to add extra [grpc.DialOption] for a specific [Address].
// It may be used to define different transport credentials for different services.
func WithAddressDialOptions(address Address, opts ...grpc.DialOption) grpc.DialOption {
	return addressDialOption{
		EmptyDialOption: grpc.EmptyDialOption{},
		address:         address,
		opts:            opts,
	}
}

type backOffDialOption struct {
	grpc.EmptyDialOption
	backOff     backoff.BackOff
	isRetriable func(error) bool
}

func defaultIsRetriable(err error) bool {
	if err == nil {
		return false
	}

	// Do not retry on context errors or if the error is definitively unrecoverable
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}

	// Assume other errors are transient and retryable
	return true
}

func WithBackOffDialOption(backOff backoff.BackOff) grpc.DialOption {
	return backOffDialOption{
		EmptyDialOption: grpc.EmptyDialOption{},
		backOff:         backOff,
		isRetriable:     defaultIsRetriable,
	}
}

func WithBackOffAndCustomRetriableDialOption(
	backOff backoff.BackOff,
	retriable func(error) bool,
) grpc.DialOption {
	return backOffDialOption{
		EmptyDialOption: grpc.EmptyDialOption{},
		backOff:         backOff,
		isRetriable:     retriable,
	}
}

type dialer struct {
	logger *slog.Logger
	opts   []grpc.DialOption
	group  singleflight.Group
	mu     sync.RWMutex
	closed bool
	cons   map[Address]*grpc.ClientConn
}

func NewDialer(logger *slog.Logger, opts ...grpc.DialOption) Dialer {
	return &dialer{
		logger: logger,
		opts:   opts,
		group:  singleflight.Group{},
		mu:     sync.RWMutex{},
		closed: false,
		cons:   map[Address]*grpc.ClientConn{},
	}
}

func (d *dialer) Dial(ctx context.Context, address Address) (grpc.ClientConnInterface, error) {
	conn, err := d.getExisting(address)
	if conn != nil || err != nil {
		return conn, err
	}

	return d.dial(ctx, address)
}

func (d *dialer) getExisting(address Address) (*grpc.ClientConn, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if d.closed {
		return nil, ErrDialerClosed
	}
	if conn, ok := d.cons[address]; ok {
		return conn, nil
	}
	return nil, nil //nolint:nilnil // don't want to introduce a sentinel error in private method
}

func (d *dialer) dial(ctx context.Context, address Address) (*grpc.ClientConn, error) { //nolint:gocognit
	log := d.logger.With(slog.String("address", string(address)))
	res, err, _ := d.group.Do(string(address), func() (interface{}, error) {
		log.DebugContext(ctx, "connecting to grpc server")
		start := time.Now()
		opts := d.opts
		for _, opt := range opts {
			if c, ok := opt.(addressDialOption); ok {
				if c.address == address {
					opts = append(opts, c.opts...)
				}
			}
		}
		// Check if backOffDialOption is specified and get backoff configuration
		var backOff backoff.BackOff
		retriable := defaultIsRetriable
		for _, opt := range opts {
			if retryOpt, ok := opt.(backOffDialOption); ok {
				backOff = retryOpt.backOff
				retriable = retryOpt.isRetriable
			}
		}

		var err error
		var conn *grpc.ClientConn
		if backOff != nil {
			backOff.Reset()
		}
		timer := time.Timer{}
		for {
			conn, err = grpc.DialContext(ctx, string(address), opts...) //nolint:staticcheck // TODO: use NewClient
			if err == nil || backOff == nil || !retriable(err) {
				break
			}

			wait := backOff.NextBackOff()
			if wait == backoff.Stop {
				return nil, err
			}
			timer.Reset(wait)

			select {
			case <-ctx.Done():
				return conn, ctx.Err()
			case <-timer.C:
			}
		}
		if err != nil {
			return nil, err
		}
		log.DebugContext(ctx, "connected to grpc server", slog.Duration("elapsed", time.Since(start)))

		d.mu.Lock()
		defer d.mu.Unlock()
		if d.closed {
			_ = conn.Close()
			return nil, ErrDialerClosed
		}
		d.cons[address] = conn

		return conn, nil
	})
	if err != nil {
		return nil, err
	}

	return res.(*grpc.ClientConn), nil
}

func (d *dialer) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.closed {
		return nil
	}
	var errs error
	for address, conn := range d.cons {
		log := d.logger.With(slog.String("address", string(address)))
		log.Debug("closing grpc connect")
		start := time.Now()
		err := conn.Close()
		if status.Code(err) == codes.Canceled {
			continue
		}
		if err != nil {
			errs = errors.Join(errs, fmt.Errorf("close connection to %s: %w", address, err))
		}
		log.Debug("grpc connection closed", slog.Duration("elapsed", time.Since(start)))
	}
	d.closed = true
	d.cons = nil
	return errs
}
