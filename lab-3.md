# Lab 3 - Up in the :cloud:

In this lab, we will look at transforming the application into a twitter bot. To complete this, you must have a twitter developer account set up with the API keys to hand.

### Step 1

If you haven't already, register for a [developer account](https://developer.twitter.com/en/docs/basics/developer-portal/overview) so you can access Twitter API keys for your application. As previously mentioned this could take a few minutes so you'll need to be patient! 🙂 

![Twitter Developer Sign up](.gitbook/assets/twitterdevacc.png)

### TODO: How to create a client app in twitter UI

### Step 2

Now that you have created an app in Twitter and your deployment platform is set up, lets quickly code it 🍻 

1. Create a new folder inside your `pkg` directory and call it `twitter_auth`
2. Inside your `twitter_auth` directory create a file called `twitter_auth.go`
3. The first thing you need to do is authenticate with twitter and connect to the app you created. To do this, read and add the following code to this file:

```go
package twitterauth

import (
    "os"

    "github.com/dghubble/go-twitter/twitter"
    "github.com/dghubble/oauth1"
    logr "github.com/sirupsen/logrus"
)

// Credentials struct contains API credentials pulled from env vars:
type Credentials struct {
    ConsumerKey       string
    ConsumerSecret    string
    AccessToken       string
    AccessTokenSecret string
}

func GetCredentials() Credentials {
    creds := Credentials{
        AccessToken:       os.Getenv("ACCESS_TOKEN"),
        AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
        ConsumerKey:       os.Getenv("CONSUMER_KEY"),
        ConsumerSecret:    os.Getenv("CONSUMER_SECRET"),
    }

    return creds
}

/* GetUserClient:
Input = credentials
Return = authenticated twitter client, error
*/
func GetUserClient(creds *Credentials) (*twitter.Client, error) {

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
```

Now that the authentication package has been created you need to call this from your `main.go` and add another route handler. To do this start by adding a new function into your `main.go`:

```go
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

    tweet, resp, err := client.Statuses.Update("Todays dad joke is: "+dadJoke, nil)
    if err != nil {
        logr.Error(err)
    }

    logr.Infof("%+v\n", resp)
    logr.Infof("%+v\n", tweet)
}
```

Once the new function has been added, in your `main()` function add the following line, just like you did before with the `jokeHandler`:

```go
r.HandleFunc("/tweetjoke", TweetHandler)
```

Because the `twitter_auth` is within its own package you will also need to add it to your imports. This will be a relative path to the file on your machine. For example mine is:

```go
twitter_auth "github.com/IBMDeveloperUK/twitter-bot-ws/pkg/twitter_auth"
```

{% hint style="info" %}
Note: It is important to keep the `twitter_auth` prefix to prevent it interfering with other declarations of the twitter API package within the code
{% endhint %}
