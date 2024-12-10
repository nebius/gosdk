# Advanced Scenarios

> For basic usage instructions, refer to the [Getting Started](/#getting-started) section.

## Logging

The Nebius SDK supports customizable logging through the `log/slog` package and a logging interceptor from the [`grpc-ecosystem` module](https://pkg.go.dev/github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging).
By default, logging is disabled.
To enable and configure logging, you can use the following options during SDK initialization:

- `gosdk.WithLogger`: Accepts a `log/slog` logger instance.
- `gosdk.WithLoggerHandler`: Accepts a handler for `log/slog`.
- `gosdk.WithLoggingOptions`: Configures advanced gRPC logging options.

### Enabling Logging

#### Using Logger

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
sdk, err := gosdk.New(ctx, gosdk.WithLogger(logger))
```

#### Using Handler

```go
handler := slog.NewJSONHandler(os.Stdout, nil)
sdk, err := gosdk.New(ctx, gosdk.WithLoggerHandler(handler))
```

### Customizing gRPC Logging

To customize gRPC logging, use the `gosdk.WithLoggingOptions` constructor option.
This option allows you to specify additional configurations for the grpc-ecosystem logging interceptor, including which events to log and other logging behaviors.

For example, to log all events (such as start, finish, payload sent, and payload received), configure the SDK like this:

```go
import "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"

sdk, err := gosdk.New(
	ctx,
	gosdk.WithLogger(logger),
	gosdk.WithLoggingOptions(
		logging.WithLogOnEvents(
			logging.StartCall,
			logging.FinishCall,
			logging.PayloadReceived,
			logging.PayloadSent,
		),
	),
)
```

The `gosdk.WithLoggingOptions` option also accepts other settings, enabling you to fine-tune the logging behavior further.
Refer to the grpc-ecosystem [documentation](https://pkg.go.dev/github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging#Option) for a full list of supported options.

### Integrating with Other Logging Libraries

Although the SDK natively supports `log/slog`, you can integrate it with other logging libraries using adapters.

#### Example: Using the `zap` Logging Library

To integrate with `zap`, you can use the `slog-zap` adapter as shown below:

```go
import "go.uber.org/zap"
import slogzap "github.com/samber/slog-zap/v2"

zapLogger, _ := zap.NewProduction()
handler := slogzap.Option{Level: slog.LevelDebug, Logger: zapLogger}.NewZapHandler()
sdk, err := gosdk.New(ctx, gosdk.WithLoggerHandler(handler))
```

This approach allows you to harness the power of your preferred logging framework while maintaining compatibility with the Nebius SDK.

## Error Handling

The SDK offers robust error-handling mechanisms to help you build resilient applications.
By default, you can wrap and propagate SDK errors, which suffices for most scenarios:

```go
operation, err := sdk.Services().Compute().V1().Instance().Create(ctx, req)
if err != nil {
	return fmt.Errorf("create instance: %w", err)
}

operation, err = operation.Wait(ctx)
if err != nil {
	return fmt.Errorf("wait for instance creation: %w", err)
}
```

However, in advanced use cases, the SDK provides tools to inspect and handle specific error types, such as gRPC errors, authentication issues, and service error details.

### Handling gRPC Errors

The SDK is built as a wrapper over the gRPC client.
Most errors include gRPC status codes, which you can inspect using standard `grpc` utilities.

> ⚠️ **Note:** This approach does not apply to authentication errors. See the next section for handling those.

#### Checking gRPC Status Code

```go
import "google.golang.org/grpc/codes"
import "google.golang.org/grpc/status"

instance, err := sdk.Services().Compute().V1().Instance().Get(ctx, req)
if err != nil {
	if status.Code(err) == codes.NotFound {
		// Handle resource not found
	}
	return fmt.Errorf("get instance: %w", err)
}
```

#### Accessing the Full gRPC Status

```go
import "google.golang.org/grpc/status"

instance, err := sdk.Services().Compute().V1().Instance().Get(ctx, req)
if err != nil {
	if st, isGrpc := status.FromError(err); isGrpc {
		// Process the gRPC status
	}
	return fmt.Errorf("get instance: %w", err)
}
```

### Handling Authentication Errors

Authentication errors in the SDK are encapsulated in the `auth.Error` type.
This type does not implement `Unwrap`, ensuring precise error handling.
To access the underlying error, use `errors.As`:

```go
import "github.com/nebius/gosdk/auth"
import "google.golang.org/grpc/status"

instance, err := sdk.Services().Compute().V1().Instance().Get(ctx, req)
if err != nil {
	var authErr *auth.Error
	if errors.As(err, &authErr) {
		underlyingErr := authErr.Cause
		if st, isGrpc := status.FromError(underlyingErr); isGrpc {
			// Handle gRPC errors from IAM authentication
		}
	}
	return fmt.Errorf("get instance: %w", err)
}
```

### Extracting Service Error Details

Many Nebius services return additional error details through the `serviceerror` package.
These details can provide insights into issues like quota failures or validation errors.

#### Adding Detailed Information to Error Messages

```go
import "github.com/nebius/gosdk/serviceerror"

instance, err := sdk.Services().Compute().V1().Instance().Get(ctx, req)
if err != nil {
	var sErr *serviceerror.Error
	if errors.As(err, &sErr) {
		return fmt.Errorf("get instance: %w\n%s", sErr, sErr.Cause())
	}
	return fmt.Errorf("get instance: %w", err)
}
```

#### Handling Specific Error Details

```go
import "github.com/nebius/gosdk/serviceerror"

operation, err := operation.Wait(ctx)
if err != nil {
	var sErr *serviceerror.Error
	if errors.As(err, &sErr) {
		for _, detail := range sErr.Details {
			failure, isQuotaFailure := detail.(*serviceerror.QuotaFailure)
			if isQuotaFailure {
				// do something special
			}
		}
	}
	return fmt.Errorf("wait: %w", err)
}
```

For further information, see the [`serviceerror` package reference](https://pkg.go.dev/github.com/nebius/gosdk/serviceerror).

### Handling Operation Errors

When the SDK successfully polls a completed operation that fails, the `Wait` method returns an `*operations.Error`.
This eliminates the need to call `operation.Successful()` manually.

```go
import "github.com/nebius/gosdk/operations"

operation, err := operation.Wait(ctx)
if err != nil {
	var opErr *operations.Error
	if errors.As(err, &opErr) {
		// Handle operation failure
	} else {
		// Handle other errors
	}
}
```

## TODO (unordered)

- Retry policy and idempotency key + how to custom idempotency key
- Trace-ID and Request-ID from the server for support
- Service account token refresh in background
- Custom gRPC dial options
- Working with gRPC call options + our custom options (conn.WithIdempotencyKey, operations.PollInterval)
- sdk.Init(), gosdk.WithInit() and gosdk.WithExplicitInit()
- How to enable tracing?
- How to enable metrics?
- Tests and mocks
- gosdk.CustomTokener usage example
- Sensitive fields sanitizing
- Everything about operations (see api repo) + methods GetOperation and ListOperations in all services + operations.PollInterval
