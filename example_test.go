package gosdk_test

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/nebius/gosdk"
	"github.com/nebius/gosdk/auth"
	compute "github.com/nebius/gosdk/proto/nebius/compute/v1"
)

// This example demonstrates how to create and configure a new [gosdk.SDK] instance with logging and credentials,
// then retrieve a list of compute instances using the gosdk API.
//
// After usage, the SDK instance should be closed.
//
// Note: Be sure to replace `creds` with appropriate credentials.
func ExampleSDK() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	var creds gosdk.Credentials // define your credentials (see other examples)

	sdk, err := gosdk.New(
		ctx,
		gosdk.WithLogger(logger),
		gosdk.WithCredentials(creds),
	)
	if err != nil {
		return // fmt.Errorf("create gosdk: %w", err)
	}
	defer func() {
		errX := sdk.Close()
		if errX != nil {
			logger.ErrorContext(ctx, "failed to close sdk", slog.Any("error", errX))
		}
	}()

	list, err := sdk.Services().Compute().V1().Instance().List(ctx, &compute.ListInstancesRequest{
		ParentId: "my-parent",
	})
	if err != nil {
		return // fmt.Errorf("list instances: %w", err)
	}

	for _, instance := range list.GetItems() {
		fmt.Println(instance.GetMetadata().GetId())
	}
}

// This examples demonstrates using the [gosdk.WithExplicitInit] option
// to separate [gosdk.SDK] construction ([gosdk.New]) from IO operations performed during [gosdk.SDK.Init].
// This allows for explicit control over initialization timing.
func ExampleSDK_Init() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	var creds gosdk.Credentials // define your credentials (see other examples)

	sdk, err := gosdk.New(
		ctx,
		gosdk.WithLogger(logger),
		gosdk.WithCredentials(creds),
		gosdk.WithExplicitInit(true),
	)
	if err != nil {
		return // fmt.Errorf("create gosdk: %w", err)
	}
	defer func() {
		errX := sdk.Close()
		if errX != nil {
			logger.ErrorContext(ctx, "failed to close sdk", slog.Any("error", errX))
		}
	}()

	err = sdk.Init(ctx)
	if err != nil {
		return // fmt.Errorf("init gosdk: %w", err)
	}
}

func ExampleCredentials_useIAMToken() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	sdk, err := gosdk.New(
		ctx,
		gosdk.WithLogger(logger),
		gosdk.WithCredentials(
			gosdk.IAMToken(os.Getenv("MY_ENV")),
		),
	)
	if err != nil {
		return // fmt.Errorf("create gosdk: %w", err)
	}
	defer func() {
		errX := sdk.Close()
		if errX != nil {
			logger.ErrorContext(ctx, "failed to close sdk", slog.Any("error", errX))
		}
	}()
}

func ExampleCredentials_useServiceAccount() {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// you can find more ways to create a service account in the auth package
	serviceAccount := auth.NewPrivateKeyFileParser(
		nil, // or you can set up your own fs, eg: `os.DirFS("."),`
		"private/key/file/path",
		"public-key-id",
		"service-account-id",
	)

	sdk, err := gosdk.New(
		ctx,
		gosdk.WithLogger(logger),
		gosdk.WithCredentials(
			gosdk.ServiceAccountReader(serviceAccount),
		),
	)
	if err != nil {
		return // fmt.Errorf("create gosdk: %w", err)
	}
	defer func() {
		errX := sdk.Close()
		if errX != nil {
			logger.ErrorContext(ctx, "failed to close sdk", slog.Any("error", errX))
		}
	}()
}
