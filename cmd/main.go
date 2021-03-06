package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	twitter_auth "github.com/IBMDeveloperUK/twitter-bot/pkg/twitter_auth"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/gorilla/mux"
	logr "github.com/sirupsen/logrus"
)

func main() {
	logr.Info("twitter-bot-v0.0.1")

	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/readiness", readinessHandler)
	r.HandleFunc("/showjoke", JokeHandler)
	r.HandleFunc("/tweetjoke", TweetHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		logr.Info("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			logr.Error(err)
			os.Exit(1)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	srv.Shutdown(ctx)

	logr.Info("Shutting down")
	os.Exit(0)
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	logr.Infof("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TweetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	dadJoke, err := getJoke()
	if err != nil {
		logr.Error(err)
		os.Exit(1)
	}
	w.Write([]byte(fmt.Sprintf("The following joke will be tweeted, %s\n", dadJoke)))

	creds := twitter_auth.GetCredentials()

	client, err := twitter_auth.GetUserClient(&creds)
	if err != nil {
		logr.Error("Error getting Twitter Client")
		logr.Error(err)
	}

	// tweet, resp, err := client.Statuses.Update("Todays dad joke is: "+dadJoke+"\nTweeted from liams-twitter-bot :)", nil)
	// if err != nil {
	// 	logr.Error(err)
	// }

	// logr.Infof("%+v\n", resp)
	// logr.Infof("%+v\n", tweet)

	/*
	**** Search ****
	 */

	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "IBM",
	})

	if err != nil {
		logr.Error(err)
	}

	logr.Infof("%+v\n", resp)
	logr.Infof("%+v\n", search)
}

func JokeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	logr.Infof("Received request to show a joke")
	dadJoke, err := getJoke()
	if err != nil {
		logr.Error(err)
	}
	w.Write([]byte(fmt.Sprintf(dadJoke)))
	logr.Info(dadJoke)
}

/* getJoke:
Input = N/A
Return = joke as a string or error
*/
func getJoke() (string, error) {
	logr.Infof("Received request for a joke")
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "text/plain")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
