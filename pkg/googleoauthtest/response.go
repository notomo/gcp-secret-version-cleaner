package googleoauthtest

import (
	"net/http"

	"github.com/jarcoal/httpmock"
)

func TokenResponse() (string, string, httpmock.Responder) {
	return http.MethodPost,
		"https://oauth2.googleapis.com/token",
		httpmock.NewStringResponder(http.StatusOK, `{
  "access_token": "XXXX.XXXXXXXXXXXXXXXXXXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "expires_in": 3599,
  "refresh_token": "1//XXXXXXXXXXXXXXXXXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX-XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
  "scope": "https://www.googleapis.com/auth/gmail.readonly",
  "token_type": "Bearer"
}`)
}
