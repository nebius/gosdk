# Nebius Go SDK <br> [![go minimal version][go-img]][go-url] [![go tested version][go-latest-img]][go-latest-url] [![CI][ci-img]][ci-url] [![License][license-img]][license-url] [![Go Reference][godoc-img]][godoc-url]

The Nebius Go SDK is a comprehensive client library for interacting with [nebius.ai](https://nebius.ai) services.
Built on gRPC, it supports all APIs defined in the [Nebius API repository](https://github.com/nebius/api).
This SDK simplifies resource management, authentication, and communication with Nebius services, making it a valuable tool for developers.

## Installation

Add the SDK to your Go project with the following command:

```bash
go get github.com/nebius/gosdk
```

To update to the latest version:

```bash
go get -u github.com/nebius/gosdk
```

## Supported Go Versions

- **Minimum Supported Version**: Go 1.22
- The SDK is regularly tested against the latest Go release.

## Getting Started

Follow this step-by-step guide to begin using the SDK.

> Skip ahead to the [Complete Example](#complete-example) or explore [Advanced Scenarios](SCENARIOS.md).

### SDK Initialization

Initialize the SDK with appropriate options:

```go
import "github.com/nebius/gosdk"

sdk, err := gosdk.New(ctx /*, option1, option2, option3 */)
if err != nil {
	return fmt.Errorf("create gosdk: %w", err)
}
defer sdk.Close()
```

The `gosdk.New` constructor initializes the SDK.
However, **authentication is required** for functionality.
Use the `gosdk.WithCredentials` option to provide credentials.

To clean up resources properly, ensure you call `Close` when finished.

Find all available options in [options.go](options.go) ([reference](https://pkg.go.dev/github.com/nebius/gosdk#Option)).

### Authentication and Credentials

Authentication is handled by passing credentials via the `gosdk.WithCredentials` option.
Commonly used credentials include `gosdk.IAMToken` and `gosdk.ServiceAccountReader`.

Find all available credentials in [credentials.go](credentials.go) ([reference](https://pkg.go.dev/github.com/nebius/gosdk#Credentials)).

#### Using an IAM Token

Supply an IAM token with `gosdk.IAMToken`. For instance, if your token is stored in an environment variable:

```go
token := os.Getenv("IAM_TOKEN")
sdk, err := gosdk.New(
	ctx,
	gosdk.WithCredentials(
		gosdk.IAMToken(token),
	),
)
```

To generate an IAM token using the `nebius` CLI ([documentation](https://docs.nebius.com/cli)):

```bash
IAM_TOKEN=$(nebius iam get-access-token)
```

#### Using a Service Account

To authenticate with a service account, provide the service account ID, public key ID, and RSA private key.
The SDK uses these details to generate a JWT and exchange it for an IAM token.
The token is automatically refreshed in the background to ensure continuous validity.

Use `gosdk.ServiceAccount` or `gosdk.ServiceAccountReader` with functions from [auth/service_account.go](auth/service_account.go).

Here are common approaches:

1. **Using a JSON Credentials File**:
   ```go
   import "github.com/nebius/gosdk/auth"
   
   sdk, err := gosdk.New(
   	ctx,
   	gosdk.WithCredentials(
   		gosdk.ServiceAccountReader(
   			auth.NewServiceAccountCredentialsFileParser(
   				nil, // nil to use the real file system
   				"~/path/to/service_account.json",
   			),
   		),
   	),
   )
   ```
   
   File format:
   ```json
   {
   	"subject-credentials": {
   		"alg": "RS256",
   		"private-key": "PKCS#8 PEM with new lines escaped as \n",
   		"kid": "public-key-id",
   		"iss": "service-account-id",
   		"sub": "service-account-id"
   	}
   }
   ```
2. **Using a PEM-Encoded Private Key File**:
   ```go
   auth.NewPrivateKeyFileParser(
   	nil, // nil to use the real file system
   	"~/path/to/private_key.pem",
   	"public-key-id",
   	"service-account-id",
   )
   ```
3. **Providing Key Content Directly**:
   ```go
   privateKey, _ := os.ReadFile("path/to/private_key.pem")
   auth.NewPrivateKeyParser(
   	privateKey,
   	"public-key-id",
   	"service-account-id",
   )
   ```

## Resources and Operations

Nebius communicates via gRPC, with read operations such as `Get` and `List` returning protobuf messages that describe resources.
Mutating operations like `Create`, `Update`, and `Delete` return an `Operation` object.

Operations can be either synchronous or asynchronous.
Synchronous operations are completed immediately, while asynchronous operations may take time to finish.
To ensure an operation is fully completed, use the `Wait` method, which polls the operation status until it is done.

> ℹ️ **Note**: If the operation fails, `Wait` will return an error.

### Create Resource

The following example demonstrates how to create a compute instance and use `Wait` to verify the operation's success.

```go
import common "github.com/nebius/gosdk/proto/nebius/common/v1"
import compute "github.com/nebius/gosdk/proto/nebius/compute/v1"

instanceService := sdk.Services().Compute().V1().Instance()

operation, err := instanceService.Create(ctx, &compute.CreateInstanceRequest{
	Metadata: &common.ResourceMetadata{
		ParentId: "my-project-id",
		Name:     "my-instance",
	},
	Spec: &compute.InstanceSpec{
		// instance configuration
	},
})
if err != nil {
	return fmt.Errorf("create instance: %w", err)
}

operation, err = operation.Wait(ctx)
if err != nil {
	return fmt.Errorf("wait for instance create: %w", err)
}

instanceID := operation.ResourceID()
```

### Get Resource

The operation doesn't contain the current state of the resource. Once it completes, fetch the resource.

```go
instance, err := instanceService.Get(ctx, &compute.GetInstanceRequest{
	Id: instanceID,
})
if err != nil {
	return fmt.Errorf("get instance: %w", err)
}
```

#### Get Resource by Name

Most resources can also be retrieved using their names.

```go
instance, err = instanceService.GetByName(ctx, &common.GetByNameRequest{
	ParentId: "my-project-id",
	Name:     "my-instance",
})
if err != nil {
	return fmt.Errorf("get instance by name: %w", err)
}
```

### Update Resource

When updating, provide a **complete resource specification**, not just the fields you wish to modify.
Treat the `Update` method as having full-replace semantics.

```go
operation, err := instanceService.Update(ctx, &compute.UpdateInstanceRequest{
	Metadata: &common.ResourceMetadata{
		Id:       instanceID,
		ParentId: "my-project-id",
		Name:     "my-instance",
	},
	Spec: &compute.InstanceSpec{
		// new configuration
	},
})
if err != nil {
	return fmt.Errorf("update instance: %w", err)
}

operation, err = operation.Wait(ctx)
if err != nil {
	return fmt.Errorf("wait for instance update: %w", err)
}
```

### List Resources

The `List` method retrieves resources within a specified container.
It supports pagination, which may require additional handling for large number of resources.

```go
list, err := instanceService.List(ctx, &compute.ListInstancesRequest{
	ParentId: "my-project-id",
})
if err != nil {
	return fmt.Errorf("list instances: %w", err)
}

for _, instance := range list.GetItems() {
	// process instance
}

if list.GetNextPageToken() != "" {
	list, err = instanceService.List(ctx, &compute.ListInstancesRequest{
		ParentId:  "my-project-id",
		PageToken: list.GetNextPageToken(),
	})
	if err != nil {
		return fmt.Errorf("list instances: %w", err)
	}

	// repeat processing
}
```

### Filter Resources

> ⚠️ Requires Go 1.23 or `GOEXPERIMENT=rangefunc` in Go 1.22.

The `Filter` method simplifies resource listing by iterating over items across pages in a single loop.

```go
req := &compute.ListInstancesRequest{ParentId: "my-project-id"}
for instance, err := range instanceService.Filter(ctx, req) {
	if err != nil {
		return fmt.Errorf("list instances: %w", err)
	}

	// process instance
}
```

### Delete Resource

The following example demonstrates how to delete a compute instance and use `Wait` to verify the operation's success.

```go
operation, err := instanceService.Delete(ctx, &compute.DeleteInstanceRequest{
	Id: instanceID,
})
if err != nil {
	return fmt.Errorf("delete instance: %w", err)
}

operation, err = operation.Wait(ctx)
if err != nil {
	return fmt.Errorf("wait for instance delete: %w", err)
}
```

## Complete Example

This example demonstrates how to initialize the SDK with IAM token authentication and perform basic resource operations.

```go
package example

import (
	"context"
	"fmt"
	"os"

	"github.com/nebius/gosdk"
	common "github.com/nebius/gosdk/proto/nebius/common/v1"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
)

func Example() error {
	ctx := context.Background()

	// Initialize SDK with IAM token
	sdk, err := gosdk.New(
		ctx,
		gosdk.WithCredentials(
			gosdk.IAMToken(os.Getenv("IAM_TOKEN")),
		),
	)
	if err != nil {
		return fmt.Errorf("create gosdk: %w", err)
	}
	defer sdk.Close()

	instanceService := sdk.Services().Compute().V1().Instance()

	// Create resource
	operation, err := instanceService.Create(ctx, &compute.CreateInstanceRequest{
		Metadata: &common.ResourceMetadata{
			ParentId: "my-project-id",
			Name:     "my-instance",
		},
		Spec: &compute.InstanceSpec{
			// configuration
		},
	})
	if err != nil {
		return fmt.Errorf("create instance: %w", err)
	}

	// Wait for the create operation to complete successfully
	operation, err = operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("wait for instance creation: %w", err)
	}

	instanceID := operation.ResourceID()

	// Get resource by ID
	instance, err := instanceService.Get(ctx, &compute.GetInstanceRequest{
		Id: instanceID,
	})
	if err != nil {
		return fmt.Errorf("get instance: %w", err)
	}

	// Get resource by name
	instance, err = instanceService.GetByName(ctx, &common.GetByNameRequest{
		ParentId: "my-project-id",
		Name:     "my-instance",
	})
	if err != nil {
		return fmt.Errorf("get instance by name: %w", err)
	}

	// Update resource
	operation, err = instanceService.Update(ctx, &compute.UpdateInstanceRequest{
		Metadata: &common.ResourceMetadata{
			Id:       instanceID,
			ParentId: "my-project-id",
			Name:     "my-instance",
		},
		Spec: &compute.InstanceSpec{
			// updated configuration
		},
	})
	if err != nil {
		return fmt.Errorf("update instance: %w", err)
	}

	// Wait for update operation complete successfully
	operation, err = operation.Wait(ctx)
	if err != nil {
		return fmt.Errorf("wait for instance update: %w", err)
	}

	// Iterate over all resources inside container
	req := &compute.ListInstancesRequest{ParentId: "my-project-id"}
	for instance, err = range instanceService.Filter(ctx, req) {
		if err != nil {
			return fmt.Errorf("list instances: %w", err)
		}

		if instance.GetMetadata().GetId() == instanceID {
			continue // skip just created instance
		}

		// Delete resource
		operation, err = instanceService.Delete(ctx, &compute.DeleteInstanceRequest{
			Id: instance.GetMetadata().GetId(),
		})
		if err != nil {
			return fmt.Errorf("delete instance: %w", err)
		}

		// Wait for delete operation complete successfully
		operation, err = operation.Wait(ctx)
		if err != nil {
			return fmt.Errorf("wait for instance delete: %w", err)
		}
	}

	return nil
}
```

## Advanced Scenarios

Explore advanced usage examples in [SCENARIOS.md](SCENARIOS.md).

## Contributing

Contributions are welcome! Please refer to the [contributing guidelines](CONTRIBUTING.md) for more information.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

Copyright (c) 2024 Nebius B.V.



[go-img]: https://img.shields.io/github/go-mod/go-version/nebius/gosdk
[go-url]: /go.mod
[go-latest-img]: https://img.shields.io/github/go-mod/go-version/nebius/gosdk?filename=.github%2Flatest-deps%2Fgo.mod&label=tested
[go-latest-url]: /.github/latest-deps/go.mod
[ci-img]: https://github.com/nebius/gosdk/actions/workflows/ci.yml/badge.svg
[ci-url]: https://github.com/nebius/gosdk/actions/workflows/ci.yml
[license-img]: https://img.shields.io/github/license/nebius/gosdk.svg
[license-url]: /LICENSE
[godoc-img]: https://pkg.go.dev/badge/github.com/nebius/gosdk.svg
[godoc-url]: https://pkg.go.dev/github.com/nebius/gosdk
