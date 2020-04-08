package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJokeHandler(t *testing.T) {
	t.Run("Returns a joke and checks the status code", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/showjoke", nil)
		response := httptest.NewRecorder()

		JokeHandler(response, request)

		got := response.Code
		want := http.StatusOK

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestTweetHandler(t *testing.T) {
	t.Run("Returns a status code from Twitter API call", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/tweetjoke", nil)
		response := httptest.NewRecorder()

		TweetHandler(response, request)

		got := response.Code
		want := http.StatusOK

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
