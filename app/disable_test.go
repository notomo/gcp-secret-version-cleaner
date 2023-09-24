package app_test

import (
	"context"
	"testing"

	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/jarcoal/httpmock"
	"github.com/notomo/gcp-secret-version-cleaner/app"
	"github.com/notomo/gcp-secret-version-cleaner/pkg/googleoauthtest"
	"github.com/notomo/gcp-secret-version-cleaner/pkg/httpmockext"
	"github.com/notomo/gcp-secret-version-cleaner/pkg/secretmanagerext/secretmanagertest"
)

func TestDisable(t *testing.T) {
	googleoauthtest.SetGoogleApplicationCredentials(t)

	projectName := "test"
	secretName := "test_secret"
	logDir := "/tmp/gcp-secret-version-cleaner-test"

	t.Run("disables secret versions", func(t *testing.T) {
		transport := httpmock.NewMockTransport()
		defer httpmockext.AssertCalled(t, transport)

		transport.RegisterResponder(googleoauthtest.TokenResponse())
		transport.RegisterResponder(secretmanagertest.ListVersionsResponse(t, projectName, secretName, []secretmanagerpb.SecretVersion_State{
			secretmanagerpb.SecretVersion_ENABLED,
			secretmanagerpb.SecretVersion_ENABLED,
			secretmanagerpb.SecretVersion_ENABLED,
		}))
		transport.RegisterResponder(secretmanagertest.DisableVersionResponse(t, projectName, secretName, 1))
		transport.RegisterResponder(secretmanagertest.DisableVersionResponse(t, projectName, secretName, 2))
		transport.RegisterResponder(secretmanagertest.DisableVersionResponse(t, projectName, secretName, 3))

		err := app.Disable(
			context.Background(),
			projectName,
			secretName,
			"",
			0,
			app.LogTransport(logDir, transport),
			false,
		)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("does not disable secret version in dry run", func(t *testing.T) {
		transport := httpmock.NewMockTransport()
		defer httpmockext.AssertCalled(t, transport)

		transport.RegisterResponder(googleoauthtest.TokenResponse())
		transport.RegisterResponder(secretmanagertest.ListVersionsResponse(t, projectName, secretName, []secretmanagerpb.SecretVersion_State{
			secretmanagerpb.SecretVersion_ENABLED,
			secretmanagerpb.SecretVersion_ENABLED,
			secretmanagerpb.SecretVersion_ENABLED,
		}))

		err := app.Disable(
			context.Background(),
			projectName,
			secretName,
			"",
			0,
			app.LogTransport(logDir, transport),
			true,
		)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("ignores disabled or destroyed secret version", func(t *testing.T) {
		transport := httpmock.NewMockTransport()
		defer httpmockext.AssertCalled(t, transport)

		transport.RegisterResponder(googleoauthtest.TokenResponse())
		transport.RegisterResponder(secretmanagertest.ListVersionsResponse(t, projectName, secretName, []secretmanagerpb.SecretVersion_State{
			secretmanagerpb.SecretVersion_DESTROYED,
			secretmanagerpb.SecretVersion_DESTROYED,
		}))

		err := app.Disable(
			context.Background(),
			projectName,
			secretName,
			"",
			0,
			app.LogTransport(logDir, transport),
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
			secretmanagerpb.SecretVersion_ENABLED,
			secretmanagerpb.SecretVersion_ENABLED,
			secretmanagerpb.SecretVersion_ENABLED,
			secretmanagerpb.SecretVersion_ENABLED,
		}))
		transport.RegisterResponder(secretmanagertest.DisableVersionResponse(t, projectName, secretName, 1))
		transport.RegisterResponder(secretmanagertest.DisableVersionResponse(t, projectName, secretName, 2))

		err := app.Disable(
			context.Background(),
			projectName,
			secretName,
			"",
			2,
			app.LogTransport(logDir, transport),
			false,
		)
		if err != nil {
			t.Error(err)
		}
	})

}
