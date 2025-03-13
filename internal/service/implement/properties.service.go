package service

import (
	"context"
	"database/sql"
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	requestDTO "ecommerce_go/internal/models/request"
	iservice "ecommerce_go/internal/service/interface"
	"ecommerce_go/internal/utils/redis"
	constant "ecommerce_go/pkg"
	"ecommerce_go/pkg/response"
	"encoding/json"
	"errors"
	"fmt"
	"time"

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
	if err == nil && result != "" {
		// if error then countinue else return
		if err := json.Unmarshal([]byte(result), &properties); err == nil {
			return properties, response.SuccessResponseCode, nil
		}
	}

	properties, err = ps.sqlc.ListPropertiesByOwner(ctx, in)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No property found for owner ID:", in.OwnerID)
			return out, response.NotFoundResponseCode, nil
		}
		fmt.Printf("Error getting properties with this owner:: %v\n", err)
		return out, response.CannotGetPropertyByOwnerCode, err
	}

	if len(properties) > 0 {
		go redis.CacheStore(ctx, constant.PrePropertiesOwner, properties[0].OwnerID, properties, 5*time.Minute)
	}

	return properties, response.SuccessResponseCode, nil
}

// GetProperty implements iservice.IPropertiesService.
func (p *PropertiesService) GetProperty(ctx context.Context, id string) (out database.GetPropertyByIDRow, code int, err error) {
	var property database.GetPropertyByIDRow
	// redis get data
	result, err := redis.GetSearchResultsFromRedis(ctx, constant.PrePropertiesId, id)
	if err == nil && result != "" {
		// if error then countinue else return
		if err := json.Unmarshal([]byte(result), &property); err == nil {
			return property, response.SuccessResponseCode, nil
		}
	}

	property, err = p.sqlc.GetPropertyByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No property found for ID:", id)
			return out, response.NotFoundResponseCode, nil
		}
		fmt.Printf("Error getting property:: %v\n", err)
		return out, response.CannotGetPropertyByOwnerCode, err
	}

	go redis.CacheStore(ctx, constant.PrePropertiesId, property.PropertyID, property, 5*time.Minute)

	return property, response.SuccessResponseCode, nil
}

// SearchProperties implements iservice.IPropertiesService.
func (ps *PropertiesService) SearchProperties(ctx context.Context, in database.SearchPropertiesParams) (out []database.SearchPropertiesRow, code int, err error) {
	var properties []database.SearchPropertiesRow
	query := fmt.Sprintf("FromPrice:%s,ToPrice%s,limit:%v,offset:%v", in.FromPrice, in.ToPrice, in.Limit, in.Offset)
	// redis get data
	result, err := redis.GetSearchResultsFromRedis(
		ctx,
		constant.PrePropertiesSearch,
		query)

	if err == nil {
		// if error then countinue else return
		if err := json.Unmarshal([]byte(result), &properties); err == nil {
			return properties, response.SuccessResponseCode, nil
		}
	}

	properties, err = ps.sqlc.SearchProperties(ctx, in)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No property was found")
			return out, response.NotFoundResponseCode, nil
		}
		fmt.Printf("Error getting properties with this owner:: %v\n", err)
		return out, response.CannotGetPropertyByOwnerCode, err
	}

	go redis.CacheStore(ctx, constant.PrePropertiesSearch, query, properties, 5*time.Minute)

	return properties, response.SuccessResponseCode, nil
}

func (ps *PropertiesService) UpdateProperty(ctx context.Context, in requestDTO.PropertyUpdateRequest, propertyID string) (out string, code int, err error) {
	var updateProperty = database.UpdatePropertyParams{
		ID:          propertyID,
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
	global.Logger.Info("Update Properties success::", zap.Any("value::", updateProperty))

	defer redis.DeleteCache(ctx, constant.PrePropertiesOwner, in.OwnerID)
	defer redis.DeleteCache(ctx, constant.PrePropertiesId, propertyID)

	return updateProperty.ID, response.SuccessResponseCode, nil
}

func (ps *PropertiesService) DeleteProperty(ctx context.Context, id string) (code int, err error) {
	err = ps.sqlc.SoftDeleteProperty(ctx, id)

	if err != nil {
		global.Logger.Error("Delete Property failed", zap.Error(err))
		return response.CannotCreatePropertyCode, err
	}
	global.Logger.Info("Delete Property success::", zap.Any("property id::", id))

	defer redis.DeleteCache(ctx, constant.PrePropertiesId, id)
	return response.SuccessResponseCode, nil
}

func NewPropertiesService(r *database.Queries) iservice.IPropertiesService {
	return &PropertiesService{sqlc: r}
}
