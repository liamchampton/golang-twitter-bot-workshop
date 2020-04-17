# Lab 2

## Lets get RESTful💃 

### Step 1

In this lab you are going to create a web app with some routes. To do this you will use the 3rd party import `gorilla/mux`. Some bedtime reading about this can be found [here](https://github.com/gorilla/mux). We will then follow this up to output a random joke by calling an open API without the need for authentication. The API in this lab is a dad joke API but feel free to explore and chose another if you'd like, the principals are the same!

```go
func handler(w http.ResponseWriter, r *http.Request) {
    name := "<your name here>"
    logr.Info("Received request for the home page")
    w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
}
```

{% hint style="info" %}
**Note**: You will also need to add the following import to your code \(this makes the terminal logs look pretty\) :smile:
{% endhint %}

```go
logr "github.com/sirupsen/logrus"
```

### Step 2

Now you have got a route handler, you need to create the web server to invoke it. To do this, use the code snippet below and insert it into your `main()` function. Instead of using the standard go `net/http` library's we will use a more powerful 3rd party import, `gorilla mux`.

```go
// Create Server and Route Handlers
    r := mux.NewRouter()
    r.HandleFunc("/", handler)

    http.Handle("/", r)
    logr.Info("Starting up on 8080")
    logr.Error(http.ListenAndServe(":8080", nil))
```

{% hint style="info" %}
**Note**: If your plugin didn't already add the gorilla mux import, add the following line of code into your imports
{% endhint %}

```go
"github.com/gorilla/mux"
```

### Step 3

Head back to your terminal window and run your code using the command `go run cmd/main.go`. This will start up a server on port :8080.

{% hint style="info" %}
**Note**: You may be prompted by your system to allow a network connection \(you need to allow this otherwise the application may not run corretly\)
{% endhint %}

Open up a browser and type `localhost:8080` into the top URL bar and you should see the output from the `handler()` function on your screen.

### Step 4

The server is up and running but this is very basic. Use `control+c` in your terminal to terminate the server connection. Next, you will add in a route to call an API without any authentication. In this workshop, the one provided calls a random joke generator but you can change this API to be whatever you'd like.

To do this, add the the code snippet below as a new function inside your `main.go` file.

```go
func getJoke() (string, error) {
    logr.Infof("Getting joke from API..")

    req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)

    // Check the request doesnt return an error
    if err != nil {
        return "", err
    }
    req.Header.Set("Accept", "text/plain")

    resp, err := http.DefaultClient.Do(req)

    // Check the request doesnt return an error
    if err != nil {
        return "", err
    }

    // Closes resp.Body at the end of the function (always do this to prevent memory leaks, even if it isn't used)
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body) // Read resp.Body until EOF

    // Check the ReadAll doesn't return an error
    if err != nil {
        return "", err
    }

    return string(body), nil
}
```

This will return a string of the body as you want to see the joke in plain text on the page and nil, since there is no error at this point. 

Now the API call is in place, the next thing you need to do is add another handler, just like you did in Step 1.

To do this, add the new function shown below and then invoke it in the `main()` function, just like you did with the previous handler.

```go
func jokeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK) // Write the status code 200
    logr.Infof("Received request to show a joke")

    dadJoke, err := getJoke()

    // Check the request doesnt return an error
    if err != nil {
        logr.Error(err)
    }

    w.Write([]byte(fmt.Sprintf(dadJoke))) // Write the joke to a byte array
    logr.Info(dadJoke)
}
```

```go
// Add this line below the existing "/" route 
r.HandleFunc("/showjoke", jokeHandler
```

If you run the code and navigate to `localhost:8080/showjoke` in your browser you should be presented with a randomly generated joke!

Now the jokes are flowing, lets get it up in the cloud. Continue to Lab 3 to see how this is done.

