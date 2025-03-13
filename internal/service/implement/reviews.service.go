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

type ReviewService struct {
	sqlc *database.Queries
}

// CreateReview implements iservice.IReviewService.
func (r *ReviewService) CreateReview(ctx context.Context, in requestDTO.ReviewCreateModel, userID string) (string, int, error) {
	var review = database.CreateReviewParams{
		ID:         uuid.New().String(),
		UserID:     userID,
		PropertyID: in.PropertyID,
		Rating:     sql.NullInt32{Int32: in.Rating},
		Comment:    sql.NullString{String: in.Comment},
	}

	result, err := r.sqlc.CreateReview(ctx, review)
	if err != nil {
		global.Logger.Error("Error creating reviews", zap.Error(err))
		return "", response.CannotCreateReviewCode, err
	}
	global.Logger.Info("new Reviews::", zap.Any("result::", result), zap.Any("value::", review))
	redis.DeleteCache(ctx, constant.PreRoomByPropertiesId, review.PropertyID)

	return review.ID, response.SuccessResponseCode, nil
}

// DeleteReview implements iservice.IReviewService.
func (r *ReviewService) DeleteReview(ctx context.Context, id string) (code int, err error) {
	err = r.sqlc.SoftDeleteReview(ctx, id)

	if err != nil {
		global.Logger.Error("Delete Review failed", zap.Error(err))
		return response.CannotDeleteReviewCode, err
	}
	global.Logger.Info("Delete Room success::", zap.Any("room id::", id))

	defer redis.DeleteCache(ctx, constant.PreRoomById, id)
	return response.SuccessResponseCode, nil
}

// GetReviewByID implements iservice.IReviewService.
func (r *ReviewService) GetReviewByID(ctx context.Context, id string) (out database.GetReviewByIDRow, code int, err error) {
	var review database.GetReviewByIDRow
	// redis get data
	result, err := redis.GetSearchResultsFromRedis(ctx, constant.PreReviewById, id)
	if err == nil {
		// if error then countinue else return
		if err := json.Unmarshal([]byte(result), &review); err == nil {
			return review, response.SuccessResponseCode, nil
		}
	}

	review, err = r.sqlc.GetReviewByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No property was found by ID:", id)
			return out, response.NotFoundResponseCode, nil
		}
		fmt.Printf("Error getting room:: %v\n", err)
		return out, response.CannotGetRoomByIDCode, err
	}

	go redis.CacheStore(ctx, constant.PreReviewById, id, review, 5*time.Minute)
	return review, response.SuccessResponseCode, nil
}

// GetReviewByPropertyID implements iservice.IReviewService.
func (r *ReviewService) GetReviewByPropertyID(ctx context.Context, in database.ListReviewsByPropertyParams) (out []database.ListReviewsByPropertyRow, code int, err error) {
	// redis get data
	result, err := redis.GetSearchResultsFromRedis(ctx, constant.PreReviewByPropertyId, in.PropertyID)
	if err == nil {
		// if error then countinue else return
		if err := json.Unmarshal([]byte(result), &out); err == nil {
			return out, response.SuccessResponseCode, nil
		}
	}

	out, err = r.sqlc.ListReviewsByProperty(ctx, in)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No property was found by ID:", in.PropertyID)
			return out, response.NotFoundResponseCode, nil
		}
		fmt.Printf("Error getting room:: %v\n", err)
		return out, response.CannotGetRoomByIDCode, err
	}

	go redis.CacheStore(ctx, constant.PreReviewByPropertyId, in.PropertyID, out, 5*time.Minute)
	return out, response.SuccessResponseCode, nil
}

// UpdateReview implements iservice.IReviewService.
func (r *ReviewService) UpdateReview(ctx context.Context, in requestDTO.ReviewUpdateModel, reviewID string) (string, int, error) {
	var updateRoom = database.UpdateReviewParams{
		ID:      reviewID,
		Rating:  sql.NullInt32{Int32: int32(in.Rating)},
		Comment: sql.NullString{String: in.Comment},
	}

	err := r.sqlc.UpdateReview(ctx, updateRoom)
	if err != nil {
		global.Logger.Error("Error updating room", zap.Error(err))
		return "", response.CannotCreatePropertyCode, err
	}
	global.Logger.Info("Update Room success::", zap.Any("value::", updateRoom))

	defer redis.DeleteCache(ctx, constant.PreRoomByPropertiesId, in.PropertyID)
	defer redis.DeleteCache(ctx, constant.PreRoomById, reviewID)

	return updateRoom.ID, response.SuccessResponseCode, nil
}

func NewReviewService(sqlc *database.Queries) iservice.IReviewService {
	return &ReviewService{sqlc: sqlc}
}
