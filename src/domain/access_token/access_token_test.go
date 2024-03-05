package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("expiration time should be 24 hours")
	}
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}
func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("brand new access token should not be expired")
	}
	if at.AccessToken != "" {
		t.Error("new access token should not have defined access token id")
	}
	if at.UserId != 0 {
		t.Error("new access token should not have an associated user id")
	}
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}
	if !at.IsExpired() {
		t.Error("empty access token should be expired by default")
	}
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	if at.IsExpired() {
		t.Error("access token expiring three hours from now should not be expired")
	}
}
