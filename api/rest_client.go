package api

import (
    "hoqu-tracker-api/sdk/http/rest"
    "sync"
    "github.com/spf13/viper"
    "net/url"
)

var (
    cl *rest.Client
    clOnce = &sync.Once{}
)

func NewClient() *rest.Client {
    url, _ := url.Parse(viper.GetString("api.endpoint"))

    return rest.NewClient(
        url,
        viper.GetBool("api.secure"),
        viper.GetString("project.name"),
        viper.GetString("api.authHeader"),
        viper.GetString("api.auth"),
    )
}

func GetClient() *rest.Client {
    clOnce.Do(func() {
        cl = NewClient()
    })

    return cl
}