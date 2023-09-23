package app

import (
	"context"
	"testing"

	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/jarcoal/httpmock"
	"github.com/notomo/gcp-secret-version-cleaner/pkg/googleoauthtest"
	"github.com/notomo/gcp-secret-version-cleaner/pkg/httpmockext"
	"github.com/notomo/gcp-secret-version-cleaner/pkg/secretmanagerext/secretmanagertest"
)

func TestDestroy(t *testing.T) {
	tmpDir := t.TempDir()

	googleoauthtest.CreateGoogleApplicationCredentials(t, tmpDir)

	projectName := "test"
	secretName := "test_secret"
	logDir := "/tmp/gcp-secret-version-cleaner-test"

	t.Run("deletes secret versions", func(t *testing.T) {
		transport := httpmock.NewMockTransport()
		defer httpmockext.AssertCalled(t, transport)

		transport.RegisterResponder(googleoauthtest.TokenResponse())
		transport.RegisterResponder(secretmanagertest.ListVersionsResponse(t, projectName, secretName, []secretmanagerpb.SecretVersion_State{
			secretmanagerpb.SecretVersion_ENABLED,
			secretmanagerpb.SecretVersion_DISABLED,
			secretmanagerpb.SecretVersion_ENABLED,
		}))
		transport.RegisterResponder(secretmanagertest.DestroyVersionResponse(t, projectName, secretName, 1))
		transport.RegisterResponder(secretmanagertest.DestroyVersionResponse(t, projectName, secretName, 2))
		transport.RegisterResponder(secretmanagertest.DestroyVersionResponse(t, projectName, secretName, 3))

		err := Destroy(
			context.Background(),
			projectName,
			secretName,
			"",
			0,
			LogTransport(logDir, transport),
			false,
		)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("does not delete secret version in dry run", func(t *testing.T) {
		transport := httpmock.NewMockTransport()
		defer httpmockext.AssertCalled(t, transport)

		transport.RegisterResponder(googleoauthtest.TokenResponse())
		transport.RegisterResponder(secretmanagertest.ListVersionsResponse(t, projectName, secretName, []secretmanagerpb.SecretVersion_State{
			secretmanagerpb.SecretVersion_ENABLED,
			secretmanagerpb.SecretVersion_DISABLED,
			secretmanagerpb.SecretVersion_ENABLED,
		}))

		err := Destroy(
			context.Background(),
			projectName,
			secretName,
			"",
			0,
			LogTransport(logDir, transport),
			true,
		)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("ignores destroyed secret version", func(t *testing.T) {
		transport := httpmock.NewMockTransport()
		defer httpmockext.AssertCalled(t, transport)

		transport.RegisterResponder(googleoauthtest.TokenResponse())
		transport.RegisterResponder(secretmanagertest.ListVersionsResponse(t, projectName, secretName, []secretmanagerpb.SecretVersion_State{
			secretmanagerpb.SecretVersion_DESTROYED,
			secretmanagerpb.SecretVersion_DESTROYED,
		}))

		err := Destroy(
			context.Background(),
			projectName,
			secretName,
			"",
			0,
			LogTransport(logDir, transport),
			false,
		)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("can keep recent versions", func(t *testing.T) {
		transport := httpmock.NewMockTransport()
		defer httpmockext.AssertCalled(t, transport)

		transport.RegisterResponder(googleoauthtest.TokenResponse())
		transport.RegisterResponder(secretmanagertest.ListVersionsResponse(t, projectName, secretName, []secretmanagerpb.SecretVersion_State{
			secretmanagerpb.SecretVersion_DISABLED,
			secretmanagerpb.SecretVersion_DISABLED,
			secretmanagerpb.SecretVersion_ENABLED,
			secretmanagerpb.SecretVersion_ENABLED,
		}))
		transport.RegisterResponder(secretmanagertest.DestroyVersionResponse(t, projectName, secretName, 1))
		transport.RegisterResponder(secretmanagertest.DestroyVersionResponse(t, projectName, secretName, 2))

		err := Destroy(
			context.Background(),
			projectName,
			secretName,
			"",
			2,
			LogTransport(logDir, transport),
			false,
		)
		if err != nil {
			t.Error(err)
		}
	})

}
