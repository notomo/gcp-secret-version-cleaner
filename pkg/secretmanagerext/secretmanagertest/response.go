package secretmanagertest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"github.com/jarcoal/httpmock"
)

const projectNumber = 888888888888

func ListVersionsResponse(
	t *testing.T,
	projectName string,
	secretName string,
	states []secretmanagerpb.SecretVersion_State,
) (string, string, httpmock.Responder) {
	summaries := []map[string]any{}
	for i, state := range states {
		version := i + 1
		summaries = append(summaries, map[string]any{
			"name":       fmt.Sprintf("projects/%d/secrets/%s/versions/%d", projectNumber, secretName, version),
			"createTime": "2001-01-01T00:00:00.000000Z",
			"state":      state,
			"replicationStatus": map[string]any{
				"automatic": map[string]any{},
			},
			"etag":                           fmt.Sprintf("\"%014d\"", version),
			"clientSpecifiedPayloadChecksum": true,
		})
	}

	var body bytes.Buffer
	encorder := json.NewEncoder(&body)
	encorder.SetIndent("", "  ")
	if err := encorder.Encode(map[string]any{
		"versions":  summaries,
		"totalSize": len(summaries),
	}); err != nil {
		t.Fatal(err)
	}

	url := fmt.Sprintf("https://secretmanager.googleapis.com/v1/projects/%s/secrets/%s/versions", projectName, secretName)
	return http.MethodGet,
		url,
		httpmock.NewStringResponder(http.StatusOK, body.String())
}

func DestroyVersionResponse(
	t *testing.T,
	projectName string,
	secretName string,
	version int,
) (string, string, httpmock.Responder) {
	var body bytes.Buffer
	encorder := json.NewEncoder(&body)
	encorder.SetIndent("", "  ")
	if err := encorder.Encode(map[string]any{
		"name":        fmt.Sprintf("projects/%d/secrets/%s/versions/%d", projectNumber, secretName, version),
		"createTime":  "2001-01-01T00:00:00.000000Z",
		"destroyTime": "2001-01-01T00:00:00.000000Z",
		"state":       3,
		"replicationStatus": map[string]any{
			"automatic": map[string]any{},
		},
		"etag":                           fmt.Sprintf("\"8888888888888%d\"", version),
		"clientSpecifiedPayloadChecksum": true,
	}); err != nil {
		t.Fatal(err)
	}

	url := fmt.Sprintf("https://secretmanager.googleapis.com/v1/projects/%d/secrets/%s/versions/%d:destroy", projectNumber, secretName, version)
	return http.MethodPost,
		url,
		httpmock.NewStringResponder(http.StatusOK, body.String())
}

func DisableVersionResponse(
	t *testing.T,
	projectName string,
	secretName string,
	version int,
) (string, string, httpmock.Responder) {
	var body bytes.Buffer
	encorder := json.NewEncoder(&body)
	encorder.SetIndent("", "  ")
	if err := encorder.Encode(map[string]any{
		"name":       fmt.Sprintf("projects/%d/secrets/%s/versions/%d", projectNumber, secretName, version),
		"createTime": "2001-01-01T00:00:00.000000Z",
		"state":      2,
		"replicationStatus": map[string]any{
			"automatic": map[string]any{},
		},
		"etag":                           fmt.Sprintf("\"8888888888888%d\"", version),
		"clientSpecifiedPayloadChecksum": true,
	}); err != nil {
		t.Fatal(err)
	}

	url := fmt.Sprintf("https://secretmanager.googleapis.com/v1/projects/%d/secrets/%s/versions/%d:disable", projectNumber, secretName, version)
	return http.MethodPost,
		url,
		httpmock.NewStringResponder(http.StatusOK, body.String())
}
