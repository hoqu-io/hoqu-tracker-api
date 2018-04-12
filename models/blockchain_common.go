package models

// status, one of 0:NotExists, 1:Created, 2:Pending, 3:Active, 4:Done, 5:Declined
// swagger:model
// example: 3
type Status uint8

type SetStatusRequest struct {
    // id of entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id     string `json:"id"`
    // example: 3
    Status Status `json:"status"`
}

type IdRequest struct {
    // an ID of the requested entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    // in: query
    Id string `json:"id"`
}

// Add Success response
//
// swagger:response
type AddSuccessResponse struct {
    // id of created entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // Ethereum transaction hash
    // example: 0x23682ef776bfb465433e8f6a6e84eab71f039f039e86933aeca20ee46d01d576
    Tx string `json:"tx"`
}

// Tx Success response
//
// swagger:response
type TxSuccessResponse struct {
    // Ethereum transaction hash
    // example: 0x23682ef776bfb465433e8f6a6e84eab71f039f039e86933aeca20ee46d01d576
    Tx string `json:"tx"`
}