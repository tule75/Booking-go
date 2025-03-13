package iservice

import (
	"context"
	"ecommerce_go/internal/database"
	requestDTO "ecommerce_go/internal/models/request"
)

type IPropertiesService interface {
	CreateProperty(ctx context.Context, in requestDTO.PropertyCreateRequest, userID string) (out string, code int, err error)
	UpdateProperty(ctx context.Context, in requestDTO.PropertyUpdateRequest, propertyID string) (out string, code int, err error)
	SearchProperties(ctx context.Context, in database.SearchPropertiesParams) (out []database.SearchPropertiesRow, code int, err error)
	GetPropertiesByOwner(ctx context.Context, in database.ListPropertiesByOwnerParams) (out []database.ListPropertiesByOwnerRow, code int, err error)
	GetProperty(ctx context.Context, id string) (out database.GetPropertyByIDRow, code int, err error)
	DeleteProperty(ctx context.Context, id string) (code int, err error)
}

var (
	Property IPropertiesService
)

func InitProperty(p IPropertiesService) {
	Property = p
}

func GetProperty() IPropertiesService {
	if Property == nil {
		panic("implement PropertyService not found for interface IProperties")
	}
	return Property
}
