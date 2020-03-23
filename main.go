package main

import (
	"io/ioutil"
	"net/http"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
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

	creds := getCredentials()

	//fmt.Printf("%+v\n", creds)

	client, err := getUserClient(&creds)
	if err != nil {
		logr.Error("Error getting Twitter Client")
		logr.Error(err)
	}

	logr.Infof("client information %v", client)

	dadJoke, err := getJoke()
	if err != nil {
		logr.Error(err)
		os.Exit(1)
	}

	logr.Info(dadJoke)

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
