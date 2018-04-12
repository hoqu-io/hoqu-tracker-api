package models

// swagger:parameters submitWidget
type SubmitWidgetParams struct {
    // in: body
    Body SubmitWidgetRequest `json:"body"`
}

type SubmitWidgetRequest struct {
    // the ID of an Ad the widget was created for
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    AdId string `json:"adId"`
    // Widget meta data in JSON
    // example: {\"phone\": \"+447411223344\", \"email\": \"test@test.com\"}
    Meta string `json:"meta"`
    // Widget form full data in JSON
    // example: {"country": "DE", "ip": "123.44.55.66"}
    Data map[string]interface{} `json:"data"`
}
