package secretmanagerext

import (
	"context"
	"fmt"
	"net/http"
	"time"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	ghttp "google.golang.org/api/transport/http"
)

func NewClient(
	ctx context.Context,
	baseTransport http.RoundTripper,
) (*secretmanager.Client, error) {
	ctx = context.WithValue(ctx, oauth2.HTTPClient, &http.Client{
		Timeout:   20 * time.Second,
		Transport: baseTransport,
	})
	credentials, err := google.FindDefaultCredentials(ctx)
	if err != nil {
		return nil, fmt.Errorf("find default google credentials: %w", err)
	}

	transport, err := ghttp.NewTransport(ctx, baseTransport, option.WithTokenSource(credentials.TokenSource))
	if err != nil {
		return nil, fmt.Errorf("new transport: %w", err)
	}

	httpClient := &http.Client{
		Timeout:   20 * time.Second,
		Transport: transport,
	}
	client, err := secretmanager.NewRESTClient(ctx, option.WithHTTPClient(httpClient))
	if err != nil {
		return nil, fmt.Errorf("new secret manager client: %w", err)
	}

	return client, nil
}
