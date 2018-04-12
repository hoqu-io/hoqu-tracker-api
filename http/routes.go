// Package classification HOQU Tracker API by GoLang.
//
// Provides all tracker-ui widgets and webmaster integrations the opportunity to manage leads through REST requests.
//
//     Schemes: http, https
//     BasePath: /
//     Version: 0.3.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Denis Degterin<d.degterin@gmail.com> https://github.com/tracker-io/tracker-geth-api
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - hash
//     - jwt
//
//     SecurityDefinitions:
//     hash:
//          type: apiKey
//          name: X-Sign
//          in: header
//     jwt:
//         type: apiKey
//         name: X-Authorization
//         in: header
//
// swagger:meta
package http

import (
    sdkHttp "hoqu-tracker-api/sdk/http"
    "github.com/gin-gonic/gin"
    "sync"
    "hoqu-tracker-api/sdk/http/middleware"
    "hoqu-tracker-api/widget"
)

var routerOnceInitializer sync.Once

func Run() error {
    return sdkHttp.Run(initRoutes)
}

func initRoutes(router *gin.Engine) {
    routerOnceInitializer.Do(func() {
        router.Use(middleware.ExecutionTime())

        widget.InitRoutes(router)
    })
}
