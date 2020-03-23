package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	logr "github.com/sirupsen/logrus"
)

// Credentials struct contains API credentials pulled from env vars:
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// init: run before main () to load environment variables from .env into the system
func init() {
	logr.Info("Loading environment variables into the system...")
	if err := godotenv.Load(); err != nil {
		logr.Error("No .env file found")
	}
}

func main() {
	logr.Info("twitter-bot-v0.0.1")

	// Create Server and Route Handlers
	r := mux.NewRouter()

	r.HandleFunc("/", handler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/readiness", readinessHandler)
	r.HandleFunc("/showjoke", jokeHandler)
	r.HandleFunc("/tweetjoke", tweetHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// Graceful Shutdown
	waitForShutdown(srv)

	creds := getCredentials()

	//fmt.Printf("%+v\n", creds)

	client, err := getUserClient(&creds)
	if err != nil {
		logr.Error("Error getting Twitter Client")
		logr.Error(err)
	}

	logr.Infof("client information %v", client)

	/*
	**** Tweet ****
	 */
	// tweet, resp, err := client.Statuses.Update("Todays dad joke is: "+dadJoke+"\nTweeted from my bot ;)", nil)
	// if err != nil {
	// 	logr.Error(err)
	// }

	// logr.Infof("%+v\n", resp)
	// logr.Infof("%+v\n", tweet)

	/*
	**** Search ****
	 */

	// search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
	// 	Query: "IBM",
	// })

	// if err != nil {
	// 	logr.Error(err)
	// }

	// logr.Infof("%+v\n", resp)
	// logr.Infof("%+v\n", search)
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

	log.Println("Shutting down")
	os.Exit(0)
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("Received request for %s\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func tweetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	dadJoke, err := getJoke()
	if err != nil {
		logr.Error(err)
		os.Exit(1)
	}
	w.Write([]byte(fmt.Sprintf("The following joke has been tweeted, %s\n", dadJoke)))
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	dadJoke, err := getJoke()
	if err != nil {
		logr.Error(err)
		os.Exit(1)
	}

	w.Write([]byte(fmt.Sprintf(dadJoke)))
	logr.Info(dadJoke)
}

/* getJoke:
Input = N/A
Return = joke as a string or error
*/
func getJoke() (string, error) {
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

/* getCredentials:
Input = N/A
Return = Populated Credentials struct
*/
func getCredentials() Credentials {
	creds := Credentials{
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
		ConsumerKey:       os.Getenv("CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
	}

	//return creds
	return creds
}

/* getUserClient:
Input = credentials
Return = authenticated twitter client, error
*/
func getUserClient(creds *Credentials) (*twitter.Client, error) {

	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		logr.Error(err)
		return nil, err
	}

	logr.Infof("User Account Info:\n%+v\n", user)
	return client, nil
}
