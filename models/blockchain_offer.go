package models

type AddOfferRequest struct {
    // the ID of a Company the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    CompanyId    string `json:"companyId"`
    // Ethereum address of the Offer payer (who will pay for leads)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    PayerAddress string `json:"payerAddress"`
    // the name of the Offer
    Name         string `json:"name"`
    // the offer full data URL
    DataUrl      string `json:"dataUrl"`
    // the cost of the leads in HQX
    // example: 25000000000000000
    Cost         string `json:"cost"`
}

// Offer Model
//
// swagger:model
type OfferData struct {
    // an ID of the entity
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    Id string `json:"id"`
    // unix timestamp
    // example: 1518957879
    CreatedAt    string `json:"createdAt"`
    // the ID of a Company the entity is linked to
    // example: a6fdb91a-149d-11e8-b642-0ed5f89f718b
    CompanyId    string `json:"companyId"`
    // Ethereum address of the Offer payer (who will pay for leads)
    // example: 0xED2F74E1fb73b775E6e35720869Ae7A7f4D755aD
    PayerAddress string `json:"payerAddress"`
    // the name of the Offer
    Name         string `json:"name"`
    // the offer full data URL
    DataUrl      string `json:"dataUrl"`
    // the cost of the leads in HQX
    // example: 25000000000000000
    Cost         string `json:"cost"`
    // example: 3
    Status       Status  `json:"status"`
}
