package iservice

import (
	"context"
	"ecommerce_go/internal/database"
	requestDTO "ecommerce_go/internal/models/request"
)

type IReviewService interface {
	CreateReview(ctx context.Context, in requestDTO.ReviewCreateModel, userID string) (string, int, error)
	UpdateReview(ctx context.Context, in requestDTO.ReviewUpdateModel, reviewID string) (string, int, error)
	GetReviewByID(ctx context.Context, id string) (database.GetReviewByIDRow, int, error)
	GetReviewByPropertyID(ctx context.Context, in database.ListReviewsByPropertyParams) ([]database.ListReviewsByPropertyRow, int, error)
	DeleteReview(ctx context.Context, id string) (int, error)
}

var (
	ReviewService IReviewService
)

func InitReview(p IReviewService) {
	ReviewService = p
}

func GetReview() IReviewService {
	if ReviewService == nil {
		panic("implement ReviewService not found for interface IReviewService")
	}
	return ReviewService
}
