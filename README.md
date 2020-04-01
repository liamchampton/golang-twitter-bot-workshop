# Hosted Golang Twitter Bot

This workshop will show you how to build a simple Golang application and then deploy it to your preferred cloud environment

## Golang Installation
### Automated Installation (Ubuntu 16.04+ & macOS only)

Use [this open source repository](https://github.com/canha/golang-tools-install-script) to install Golang onto your machine.

Ubuntu 16.04+
```bash
wget -q -O - https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh \
| bash -s -- --version 1.14.1
```
macOS
```bash
curl https://raw.githubusercontent.com/canha/golang-tools-install-script/master/goinstall.sh \
| bash -s -- --version 1.14.1
```

Once the installation has finished, create a folder called `github.com` inside `$HOME/go/src`. Copy and paste the following command to do this for you.
```bash
mkdir $HOME/go/src/github.com
```

### Manual Installation
1. To manually install the Go tools, use the Go documentation and follow the instructions [here](https://golang.org/doc/install) 
2. Ensure your system follows the folder tree below
```bash
.
├── $HOME
│   ├── /go
|        ├── /bin
|        ├── /pkg
│        └── /src
|             └── /github.com
```
## Prerequisites
1. Twitter API Keys - Obtained by having a developer account on Twitter. This can be registered [here](https://developer.twitter.com/en/docs/basics/developer-portal/overview). This is only needed for the final lab and takes ~15minutes to set up
2. AN IDE installed (GoLand/Visual Studio Code etc)

## Lab 0 - Install the Prerequisites
### IBM cloud command line interface
(If you already have these setup and installed, go straight to [Lab 1](../Lab1/README.md))

1. Install the [IBM Cloud command-line interface from this link](https://cloud.ibm.com/docs/cli?topic=cloud-cli-install-ibmcloud-cli).     
Once installed, you can access IBM Cloud from your command-line with the prefix `ibmcloud`.
2. Log in to the IBM Cloud CLI: `ibmcloud login`.
3. Enter your IBM Cloud credentials when prompted.

   **Note:** If you have a federated ID, use `ibmcloud login --sso` to log in to the IBM Cloud CLI. Enter your user name, and use the provided URL in your CLI output to retrieve your one-time passcode. You know you have a federated ID when the login fails without the `--sso` and succeeds with the `--sso` option.

Once you are all set up you can move straight on to [Lab 1](../Lab1/README.md)


## Lab 1 - Creating a basic Golang Application :books:

### Step 1

Clone [this]() repository into `$HOME/<user>/go/github.com` and then open the `<PROJECT NAME HERE>` directory into your preferred editor. (I use Visual Studio Code with [this](https://code.visualstudio.com/docs/languages/go) recommended Go extension installed from the marketplace)

### Step 2

In a new terminal window make sure you can run the `main.go` file (located in `<PROJECTNAME>/cmd`). To do this use the command `go run cmd/main.go`. This will compile the code and run the program without building a binary (more on this later). The output should be `Hello Fellow Gopher!`.

Now the code is running successfully, you can see everything has been set up correctly and you are able to run Go code on your machine.

In the next lab we will turn this up a notch and turn our simple `hello world` program into a web server.

## Lab 2 - Lets get RESTful :dancer:

In this lab we are going to create an API without the need for authentication and add it to a route that you created in Lab 1. The API in this lab is a dad joke API but feel free to explore and chose another if you'd like!

## Lab 3 - Up in the :cloud:

Here you will have 2 options when deploying your application into a cloud environment. Chose your preferred method, or do both?

### Option 1 - IBM Cloud Foundary

### Option 2 - Kubernetes in IBM cloud

## Lab 4 - Tweet Tweet! :bird:

In this lab, we will look at transforming application into a twitter bot. To complete this, you must have a twitter developer account set up with the API keys to hand. 
