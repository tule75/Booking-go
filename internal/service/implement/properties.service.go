package service

import (
	"context"
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	requestDTO "ecommerce_go/internal/models/request"
	iservice "ecommerce_go/internal/service/interface"
	"ecommerce_go/internal/utils/redis"
	constant "ecommerce_go/pkg"
	"ecommerce_go/pkg/response"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type PropertiesService struct {
	sqlc *database.Queries
}

// CreateProperty implements iservice.IPropertiesService.
func (ps *PropertiesService) CreateProperty(ctx context.Context, in requestDTO.PropertyCreateRequest, userID string) (out string, code int, err error) {
	var newProperty = database.CreatePropertyParams{
		ID:          uuid.New().String(),
		OwnerID:     userID,
		Name:        in.Name,
		Description: in.Description,
		Location:    in.Location,
		Price:       in.Price,
		Amenities:   in.Amenities,
	}

	result, err := ps.sqlc.CreateProperty(ctx, newProperty)
	if err != nil {
		global.Logger.Error("Error creating property", zap.Error(err))
		return "", response.CannotCreatePropertyCode, err
	}
	global.Logger.Info("new User::", zap.Any("result::", result), zap.Any("value::", newProperty))

	defer redis.DeleteCache(ctx, constant.PrePropertiesOwner, userID)
	return newProperty.ID, response.SuccessResponseCode, nil

}

// GetPropertiesByOwner implements iservice.IPropertiesService.
func (ps *PropertiesService) GetPropertiesByOwner(ctx context.Context, in database.ListPropertiesByOwnerParams) (out []database.ListPropertiesByOwnerRow, code int, err error) {
	var properties []database.ListPropertiesByOwnerRow
	// redis get data
	result, err := redis.GetSearchResultsFromRedis(ctx, constant.PrePropertiesOwner, in.OwnerID)
	if err == nil {
		// if error then countinue else return
		if err := json.Unmarshal([]byte(result), &properties); err == nil {
			return properties, response.SuccessResponseCode, nil
		}
	}

	properties, err = ps.sqlc.ListPropertiesByOwner(ctx, in)

	if err != nil {
		fmt.Printf("Error getting properties with this owner:: %v\n", err)
		return out, response.CannotGetPropertyByOwnerCode, err
	}

	return properties, response.SuccessResponseCode, nil
}

// GetProperty implements iservice.IPropertiesService.
func (p *PropertiesService) GetProperty(ctx context.Context, id string) (out database.GetPropertyByIDRow, code int, err error) {
	var property database.GetPropertyByIDRow
	// redis get data
	result, err := redis.GetSearchResultsFromRedis(ctx, constant.PrePropertiesId, id)
	if err == nil {
		// if error then countinue else return
		if err := json.Unmarshal([]byte(result), &property); err == nil {
			return property, response.SuccessResponseCode, nil
		}
	}

	property, err = p.sqlc.GetPropertyByID(ctx, id)
	if err != nil {
		fmt.Printf("Error getting property:: %v\n", err)
		return out, response.CannotGetPropertyByOwnerCode, err
	}

	return property, response.SuccessResponseCode, nil
}

// SearchProperties implements iservice.IPropertiesService.
func (ps *PropertiesService) SearchProperties(ctx context.Context, in database.SearchPropertiesParams) (out []database.SearchPropertiesRow, code int, err error) {
	var properties []database.SearchPropertiesRow
	// redis get data
	result, err := redis.GetSearchResultsFromRedis(
		ctx,
		constant.PrePropertiesSearch,
		fmt.Sprintf("FromPrice:%s,ToPrice%s,limit:%v,offset:%v", in.FromPrice, in.ToPrice, in.Limit, in.Offset))

	if err == nil {
		// if error then countinue else return
		if err := json.Unmarshal([]byte(result), &properties); err == nil {
			return properties, response.SuccessResponseCode, nil
		}
	}

	properties, err = ps.sqlc.SearchProperties(ctx, in)

	if err != nil {
		fmt.Printf("Error getting properties with this owner:: %v\n", err)
		return out, response.CannotGetPropertyByOwnerCode, err
	}

	return properties, response.SuccessResponseCode, nil
}

func (ps *PropertiesService) UpdateProperty(ctx context.Context, in requestDTO.PropertyUpdateRequest) (out string, code int, err error) {
	var updateProperty = database.UpdatePropertyParams{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Location:    in.Location,
		Price:       in.Price,
		Amenities:   in.Amenities,
	}

	err = ps.sqlc.UpdateProperty(ctx, updateProperty)
	if err != nil {
		global.Logger.Error("Error creating property", zap.Error(err))
		return "", response.CannotCreatePropertyCode, err
	}
	global.Logger.Info("Update User success::", zap.Any("value::", updateProperty))

	defer redis.DeleteCache(ctx, constant.PrePropertiesOwner, in.OwnerID)
	defer redis.DeleteCache(ctx, constant.PrePropertiesId, in.ID)

	return updateProperty.ID, response.SuccessResponseCode, nil
}

func NewPropertiesService(r *database.Queries) iservice.IPropertiesService {
	return &PropertiesService{sqlc: r}
}
