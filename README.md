# HOQU Tracker API by GoLang

Provides all tracker-ui widgets and webmaster integrations the opportunity to manage leads through REST requests.

## Installation

Install [GoLang](https://www.goinggo.net/2016/05/installing-go-and-your-workspace.html) and setup local environment

### Install govendor

```console
go get -u github.com/kardianos/govendor
```

Sync vendor packages

```console
govendor sync
```

## Building

Build Tracker API

```console
cd $GOPATH/src/hoqu-tracker-api
go build
```

## Execution

### test node via rinkeby.infura.io

Run API

```console
./hoqu-tracker-api
```

### production with own node

Run API

```console
APPLICATION_ENV=pro_net ./hoqu-tracker-api
```

## Development

### Generate API documentation

Install [go-swagger](https://github.com/go-swagger/go-swagger)

Generate swagger.json

```console
cd $GOPATH/src/hoqu-tracker-api
swagger generate spec -o ./data/docs/swagger.json
```
