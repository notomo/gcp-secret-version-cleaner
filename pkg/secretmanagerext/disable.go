package secretmanagerext

import (
	"context"
	"fmt"
	"log/slog"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

func DisableVersion(
	ctx context.Context,
	client *secretmanager.Client,
	version *secretmanagerpb.SecretVersion,
	dryRun bool,
) error {
	logger := slog.Default().With("versionName", version.Name, "dryRun", dryRun)

	if version.State == secretmanagerpb.SecretVersion_DISABLED ||
		version.State == secretmanagerpb.SecretVersion_DESTROYED {
		logger.Info("already disabled or destroyed")
		return nil
	}

	logger.Info("disabling")
	if dryRun {
		return nil
	}

	if _, err := client.DisableSecretVersion(ctx, &secretmanagerpb.DisableSecretVersionRequest{
		Name: version.Name,
	}); err != nil {
		return fmt.Errorf("disable secret version: %w", err)
	}

	logger.Info("disabled")

	return nil
}
