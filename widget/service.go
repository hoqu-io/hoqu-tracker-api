package widget

import (
    "sync"
    "hoqu-tracker-api/api"
    "hoqu-tracker-api/models"
    "time"
)

var (
    serv *Service
    once = &sync.Once{}
)

type Service struct {
    leadRep *api.LeadRepository
    adRep *api.AdRepository
    offerRep *api.OfferRepository
}

func NewService() *Service {
    return &Service{
        leadRep: api.GetLeadRepository(),
        adRep: api.GetAdRepository(),
        offerRep: api.GetOfferRepository(),
    }
}

func GetService() *Service {
    once.Do(func() {
        serv = NewService()
    })

    return serv
}

func (s *Service) SubmitWidget(entity *models.SubmitWidgetRequest) (*models.AddSuccessResponse, error) {
    ad, err := s.adRep.GetById(entity.AdId)
    if err != nil {
        return nil, err
    }

    offer, err := s.offerRep.GetById(ad.OfferId)
    if err != nil {
        return nil, err
    }

    // TODO pack lead data and save to secure place
    dataUrl := "https://account.hoqu.io/tx/f57655d7f76e557a7bb6"

    lead := &models.AddLeadRequest{
        AdId: ad.Id,
        Meta: entity.Meta,
        DataUrl: dataUrl,
        Price: offer.Cost,
    }

    // TODO move price setting and lead data saving to selling stage
    resp, err := s.leadRep.CreateLead(lead)
    if err != nil {
        return nil, err
    }

    defer func(){
        go func(){
            time.Sleep(time.Duration(60) * time.Second)
            s.leadRep.SellLead(resp.Id)
        }()
    }()

    return resp, nil
}