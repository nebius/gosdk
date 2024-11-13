# Nebius Go SDK <br> [![go minimal version][go-img]][go-url] [![go tested version][go-latest-img]][go-latest-url] [![CI][ci-img]][ci-url] [![License][license-img]][license-url] [![Go Reference][godoc-img]][godoc-url]

The Nebius Go SDK is a client library for interacting with [nebius.ai](https://nebius.ai) services.
It leverages gRPC to communicate with the APIs defined in the [Nebius API repository](https://github.com/nebius/api).

## Installation

Use `go get` to retrieve the SDK to add it to your project's Go module dependencies.
```bash
go get github.com/nebius/gosdk
```

To update the SDK use `go get -u` to retrieve the latest version of the SDK.
```bash
go get -u github.com/nebius/gosdk
```

## Go Versions Supported

- The minimal supported version is Go 1.22.
- The SDK is also tested with the latest Go release.

## Usage

### Comprehensive Example

This example demonstrates how to initialize the SDK, create a new client, and use it to create a compute instance.

```go
package example

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/nebius/gosdk"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
)

func CreateComputeInstanceExample() error {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Define your credentials (see Authentication section for details)
	var creds gosdk.Credentials

	// Initialize the SDK
	sdk, err := gosdk.New(
		ctx,
		gosdk.WithLogger(logger),
		gosdk.WithCredentials(creds),
	)
	if err != nil {
		return fmt.Errorf("create gosdk: %w", err)
	}
	defer sdk.Close()

	// Create a compute instance
	operation, err := sdk.Services().Compute().V1().Instance().Create(ctx, &compute.CreateInstancesRequest{
		ParentId: "my-parent",
	})
	if err != nil {
		return fmt.Errorf("create instance: %w", err)
	}

	// Wait for the operation to complete successfully
	operation, err = operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("wait for create instance: %w", err)
	}

	instanceID := operation.ResourceID()
	fmt.Printf("Compute instance is created successfully with ID %s\n", instanceID)

	return nil
}
```

### Authentication

TODO: describe options to provide credentials.

### Async Operations

TODO: describe difference between get/list and create/update/delete and how to work with operations.

### Error Handling

TODO: describe working with serviceerrors package.

### Retry Policy

TBD

### Full Replace on Update

TBD

### Testing

TODO: describe how to use mocks to test the user's code

## Contributing

Contributions are welcome! Please refer to the [contributing guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.



[go-img]: https://img.shields.io/github/go-mod/go-version/nebius/gosdk
[go-url]: /go.mod
[go-latest-img]: https://img.shields.io/github/go-mod/go-version/nebius/gosdk?filename=.github%2Flatest-deps%2Fgo.mod&label=tested
[go-latest-url]: /.github/latest-deps/go.mod
[ci-img]: https://github.com/nebius/gosdk/actions/workflows/ci.yaml/badge.svg
[ci-url]: https://github.com/nebius/gosdk/actions/workflows/ci.yaml
[license-img]: https://img.shields.io/github/license/nebius/gosdk.svg
[license-url]: /LICENSE
[godoc-img]: https://pkg.go.dev/badge/github.com/nebius/gosdk.svg
[godoc-url]: https://pkg.go.dev/github.com/nebius/gosdk
