# exchange-rates
Use [Abstract API](https://app.abstractapi.com/) to obtain currency exchange rates.

## Setup
The application requires a file called `.env` in the current directory with an API key from Abstract to allow access to their exchange rate currency API.

Once the key has been obtained from the [exchange rate API](https://app.abstractapi.com/api/exchange-rates) page, place it in the file thus:
```
api_key: <key goes here>
```

An example file is present in the repo called `env`. Add the key to the `api_key:` line and rename the file to `.env` and it should be ready to use.

## Usage
Launch the application with two command line arguments, a base currency (e.g. USD) and the currency you want the exchange rate for (e.g. EUR).

> Note the three-letter currency codes must be used here. See the [list of supported currency codes](https://docs.abstractapi.com/exchange-rates#currency-codes-of-supported-currencies)

To run from source:

```shell
go run . GBP USD
```

## Building
The application can be compiled to an executable using standard Go build instructions, e.g.:
```shell
go build .
go build -ldflags "-s -w" .
```

### Cross-compiling examples
Building for Windows from Linux
```shell
GOOS=windows GOARCH=amd64 go build -o exchange-rate.exe .
```

Building for Linux from Windows
```shell
@echo off
set GOOS=linux
set GOARCH=amd64
go build -o exchange-rate .
```