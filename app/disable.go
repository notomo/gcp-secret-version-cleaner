package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/notomo/gcp-secret-version-cleaner/pkg/secretmanagerext"
)

func Disable(
	ctx context.Context,
	projectName string,
	secretName string,
	versionFilter string,
	keepRecentCount uint,
	baseTransport http.RoundTripper,
	dryRun bool,
) error {
	client, err := secretmanagerext.NewClient(ctx, baseTransport)
	if err != nil {
		return fmt.Errorf("new secret manager client: %w", err)
	}
	defer client.Close()

	versions, err := secretmanagerext.ListVersions(
		ctx,
		client,
		projectName,
		secretName,
		versionFilter,
		keepRecentCount,
	)
	if err != nil {
		return fmt.Errorf("list secret version names: %w", err)
	}

	for _, version := range versions {
		version := version
		if err := secretmanagerext.DisableVersion(
			ctx,
			client,
			version,
			dryRun,
		); err != nil {
			return fmt.Errorf("disable version: %w", err)
		}
	}

	return nil
}
