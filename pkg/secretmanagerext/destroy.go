package secretmanagerext

import (
	"context"
	"fmt"
	"log/slog"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func DestroyVersion(
	ctx context.Context,
	client *secretmanager.Client,
	version *secretmanagerpb.SecretVersion,
	dryRun bool,
) error {
	logger := slog.Default().With("versionName", version.Name, "dryRun", dryRun)

	if version.State == secretmanagerpb.SecretVersion_DESTROYED {
		logger.Info("already destroyed")
		return nil
	}

	logger.Info("destroying")
	if dryRun {
		return nil
	}

	if _, err := client.DestroySecretVersion(ctx, &secretmanagerpb.DestroySecretVersionRequest{
		Name: version.Name,
	}); err != nil {
		return fmt.Errorf("destroy secret version: %w", err)
	}

	logger.Info("destroyed")

	return nil
}
