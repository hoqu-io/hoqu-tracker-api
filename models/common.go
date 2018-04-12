package models

// Success
//
// swagger:response
type SuccessResponse struct {
    // in: body
    Body struct {
        Data struct {
            Success bool `json:"success"`
        } `json:"data"`
    }
}
