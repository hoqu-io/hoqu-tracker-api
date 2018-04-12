package widget

import (
    "github.com/gin-gonic/gin"
    "hoqu-tracker-api/sdk/http/rest"
    "hoqu-tracker-api/models"
)

func InitRoutes(router *gin.Engine) {
    eth := router.Group("/widget")
    {
        eth.POST("/submit", postSubmitWidgetAction)
    }
}

// swagger:route POST /widget/submit widgets submitWidget
//
// Submit widget. Create Lead on success and push it to the blockchain.
//
// Consumes:
// - application/json
// Produces:
// - application/json
// Responses:
//   200: SuccessResponse
//   400: RestErrorResponse
//
func postSubmitWidgetAction(c *gin.Context) {
    request := &models.SubmitWidgetRequest{}
    err := c.BindJSON(request)
    if err != nil {
        rest.NewResponder(c).ErrorValidation(err.Error())
        return
    }

    _, err = GetService().SubmitWidget(request)
    if err != nil {
        rest.NewResponder(c).Error(err.Error())
        return
    }

    rest.NewResponder(c).Success(gin.H{
        "success": true,
    })
}

