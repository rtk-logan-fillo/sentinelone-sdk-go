# SentinelOne SDK for Go
This project is an unoffical SentinelOne SDK for Go. It is based off of the swagger schema `./schema/sentinelone_original.json` which was sneakily acquired through SentinelOne's management console. This schema was then very heavily doctored into `./schema/sentinelone.yaml` which can be used to generate a Go client using `openapi-generate`. The new schema includes only the endpoints required at the moment, but can be extended by carefully copying over attributes from `./schema/sentinelone_original.yaml`.
## Getting Started
### Authentication
You need an API token to make calls to S1. You can get an API Token by going to the S1 dashboard and clicking in the top right corner: `<YOUR_NAME> (Admin) > My User > Options > Generate API Token`. Remember to save the token locally.
### Installing 
To install the SDK client, install it with `go get`. Note that we are installing the internal generated package and not the meta package in which this document exists.
```bash
go get github.com/rtk-logan-fillo/sentinelone-sdk-go/sentinelone
```
## Usage 
### Configuring client
You need to configure the client before using it, below is an example of how to do this. The `example` directory contains practical examples of how to use the configured client to make calls to S1 endpoints.
```go
package main

import "github.com/rtk-logan-fillo/sentinelone-sdk-go/sentinelone"

func getConfig() *sentinelone.Configuration {
	token := os.Getenv("SENTINELONE_API_TOKEN")
	if token == "" {
		panic("you must run \"export SENTINELONE_API_TOKEN=<A valid API token>\"")
	}
	cfg := sentinelone.NewConfiguration()
	cfg.Host = "usea1-partners.sentinelone.net"
	cfg.UserAgent = "Arctic-Wolf/1.0"
	cfg.Scheme = "https"
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("APIToken %s", token))
	return cfg
}

func main() {
	client := sentinelone.NewAPIClient(getConfig())
    	// ...
}
```
 Note that we pipe the API token into the config using an environment variable, so if you configure your client like this, you must run the following before running any code.
```bash
export SENTINELONE_API_TOKEN=<API_TOKEN>
```
### Running examples
To run the examples, clone this repo and install dependencies, making sure that you have exported an environment variable `SENTINELONE_API_TOKEN` in your current shell.
```bash
export SENTINELONE_API_TOKEN=<API_TOKEN>
git clone github.com/rtk-logan-fillo/sentinelone-sdk-go
cd sentinelone-sdk-go
make install
go run example/agents/main.go
```

## Developer Guide
### Install dependencies
Installs all Go dependencies from the root module and the generated `sentinelone` module
```
make install
```
### Validate Swagger schema
Use this command to validate any manual changes done to the Swagger schema. Note that generation will automatically perform this step prior to generation
```
make validate
```
### Generate API client from Swagger schema
Generates the Go client for the API based on the Swagger schema `./schema/sentinelone.yaml` into the `sentinelone` directory
```
make generate
```

## Notes
- `openapi-generator` will soon generate test files for Go generation (https://github.com/OpenAPITools/openapi-generator/pull/13560), we probably will want this feature.
- There is currently a bug in `openapi-generator` where enums are not being generated in Go code for inline enums in the schema. The SentinelOne schema uses inline enums exclusively and it's unwise to mess too much with the schema. we will want this feature when the following PR is merged and the released (https://github.com/OpenAPITools/openapi-generator/pull/3162) 
