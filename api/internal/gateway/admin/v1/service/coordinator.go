package service

import (
	"github.com/and-period/furumaru/api/internal/gateway/admin/v1/response"
	"github.com/and-period/furumaru/api/internal/user/entity"
)

type Coordinator struct {
	response.Coordinator
}

type Coordinators []*Coordinator

func NewCoordinator(coordinator *entity.Coordinator) *Coordinator {
	return &Coordinator{
		Coordinator: response.Coordinator{
			ID:                coordinator.ID,
			Status:            coordinator.Status,
			Lastname:          coordinator.Lastname,
			Firstname:         coordinator.Firstname,
			LastnameKana:      coordinator.LastnameKana,
			FirstnameKana:     coordinator.FirstnameKana,
			MarcheName:        coordinator.MarcheName,
			Username:          coordinator.Username,
			Profile:           coordinator.Profile,
			ProductTypeIDs:    coordinator.ProductTypeIDs,
			ThumbnailURL:      coordinator.ThumbnailURL,
			Thumbnails:        NewImages(coordinator.Thumbnails).Response(),
			HeaderURL:         coordinator.HeaderURL,
			Headers:           NewImages(coordinator.Headers).Response(),
			PromotionVideoURL: coordinator.PromotionVideoURL,
			BonusVideoURL:     coordinator.BonusVideoURL,
			InstagramID:       coordinator.InstagramID,
			FacebookID:        coordinator.FacebookID,
			Email:             coordinator.Email,
			PhoneNumber:       coordinator.PhoneNumber,
			PostalCode:        coordinator.PostalCode,
			Prefecture:        coordinator.Prefecture,
			City:              coordinator.City,
			AddressLine1:      coordinator.AddressLine1,
			AddressLine2:      coordinator.AddressLine2,
			CreatedAt:         coordinator.CreatedAt.Unix(),
			UpdatedAt:         coordinator.CreatedAt.Unix(),
		},
	}
}

func (c *Coordinator) Response() *response.Coordinator {
	return &c.Coordinator
}

func NewCoordinators(coordinators entity.Coordinators) Coordinators {
	res := make(Coordinators, len(coordinators))
	for i := range coordinators {
		res[i] = NewCoordinator(coordinators[i])
	}
	return res
}

func (cs Coordinators) Map() map[string]*Coordinator {
	res := make(map[string]*Coordinator, len(cs))
	for _, c := range cs {
		res[c.ID] = c
	}
	return res
}

func (cs Coordinators) Response() []*response.Coordinator {
	res := make([]*response.Coordinator, len(cs))
	for i := range cs {
		res[i] = cs[i].Response()
	}
	return res
}
