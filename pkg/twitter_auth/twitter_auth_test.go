package twitterauth

import (
	"net/http"
	"testing"
)

func TestGetCredentials(t *testing.T) {
	t.Run("Check auth credentials exist", func(t *testing.T) {

		got := GetCredentials()
		want := http.StatusOK

		if got == (Credentials{}) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
