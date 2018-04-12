package api

import (
    "hoqu-tracker-api/sdk/http/rest"
    "sync"
    "hoqu-tracker-api/models"
    "net/http"
)

var (
    orep *OfferRepository
    oonce = &sync.Once{}
)

type OfferRepository struct {
    client *rest.Client
}

func NewOfferRepository() *OfferRepository {
    return &OfferRepository{
        client: GetClient(),
    }
}

func GetOfferRepository() *OfferRepository {
    oonce.Do(func() {
        orep = NewOfferRepository()
    })

    return orep
}

func (r *OfferRepository) GetById(id string) (*models.OfferData, error) {
    entity := &models.OfferData{}
    entity.Id = id

    resp := &struct{
        Offer *models.OfferData `json:"Offer"`
    }{
        Offer: entity,
    }

    _, err := r.client.DoRest(http.MethodGet, "/platform/offer/" + id, nil, resp)
    if err != nil {
        return nil, err
    }

    return entity, nil
}
