package api

import (
    "hoqu-tracker-api/sdk/http/rest"
    "sync"
    "hoqu-tracker-api/models"
    "net/http"
)

var (
    arep *AdRepository
    aonce = &sync.Once{}
)

type AdRepository struct {
    client *rest.Client
}

func NewAdRepository() *AdRepository {
    return &AdRepository{
        client: GetClient(),
    }
}

func GetAdRepository() *AdRepository {
    aonce.Do(func() {
        arep = NewAdRepository()
    })

    return arep
}

func (r *AdRepository) GetById(id string) (*models.AdData, error) {
    entity := &models.AdData{}
    entity.Id = id

    resp := &struct{
        Ad *models.AdData `json:"Ad"`
    }{
        Ad: entity,
    }

    _, err := r.client.DoRest(http.MethodGet, "/platform/ad/" + id, nil, resp)
    if err != nil {
        return nil, err
    }

    return entity, nil
}
