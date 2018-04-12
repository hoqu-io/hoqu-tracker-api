package api

import (
    "hoqu-tracker-api/sdk/http/rest"
    "sync"
    "github.com/spf13/viper"
    "hoqu-tracker-api/models"
    "net/http"
)

var (
    lrep *LeadRepository
    lonce = &sync.Once{}
)

type LeadRepository struct {
    client *rest.Client
    trackerId string
}

func NewLeadRepository() *LeadRepository {
    return &LeadRepository{
        client: GetClient(),
        trackerId: viper.GetString("tracker.id"),
    }
}

func GetLeadRepository() *LeadRepository {
    lonce.Do(func() {
        lrep = NewLeadRepository()
    })

    return lrep
}

func (r *LeadRepository) CreateLead(entity *models.AddLeadRequest) (*models.AddSuccessResponse, error) {
    resp := &models.AddSuccessResponse{}
    entity.TrackerId = r.trackerId

    _, err := r.client.DoRest(http.MethodPost, "/platform/lead/add", entity, resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}

func (r *LeadRepository) SellLead(id string) (*models.TxSuccessResponse, error) {
    resp := &models.TxSuccessResponse{}
    entity := models.IdRequest{
        Id: id,
    }

    _, err := r.client.DoRest(http.MethodPost, "/platform/lead/sell", entity, resp)
    if err != nil {
        return nil, err
    }

    return resp, nil
}