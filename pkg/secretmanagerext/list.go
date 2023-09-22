package secretmanagerext

import (
	"cmp"
	"context"
	"errors"
	"fmt"

	"slices"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"google.golang.org/api/iterator"
)

func ListVersions(
	ctx context.Context,
	client *secretmanager.Client,
	projectName string,
	secretName string,
	versionFilter string,
	excludeRecentCount uint,
) ([]*secretmanagerpb.SecretVersion, error) {
	req := &secretmanagerpb.ListSecretVersionsRequest{
		Parent:   fmt.Sprintf("projects/%s/secrets/%s", projectName, secretName),
		Filter:   versionFilter,
		PageSize: 25000,
	}

	versions := []*secretmanagerpb.SecretVersion{}
	iter := client.ListSecretVersions(ctx, req)
	for {
		version, err := iter.Next()
		if errors.Is(err, iterator.Done) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("iter next: %w", err)
		}
		versions = append(versions, version)
	}

	slices.SortFunc(versions, func(a, b *secretmanagerpb.SecretVersion) int {
		return cmp.Compare(a.Name, b.Name)
	})

	if len(versions) < int(excludeRecentCount) {
		return versions, nil
	}
	return versions[:len(versions)-int(excludeRecentCount)], nil
}
