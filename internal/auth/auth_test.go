package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	token := "ApiKey mytoken"
	header := http.Header{"Authorization": {token}}
	headerToken, err := GetAPIKey(header)
	if err != nil {
		t.Errorf("token format is wrong: %v", err)
	}

	if headerToken != "mytoken" {
		t.Errorf("found tokens are not equal, %s != %s", headerToken, token)
	}
}
