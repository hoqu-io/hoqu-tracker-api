package main

import (
    "os"
    "hoqu-tracker-api/http"
    "hoqu-tracker-api/sdk/app"
    httpSdk "hoqu-tracker-api/sdk/http"
)

func init() {
    app.InitConfig()
    app.InitLogger()
}

func main() {
    interrupt := make(chan os.Signal, 1)

    go http.Run()

    httpSdk.StartTicking(interrupt)
}
