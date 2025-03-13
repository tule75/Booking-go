package controller

import (
	"ecommerce_go/internal/database"
	requestDTO "ecommerce_go/internal/models/request"
	iservice "ecommerce_go/internal/service/interface"
	"ecommerce_go/internal/utils/auth"
	"ecommerce_go/internal/utils/query"
	"ecommerce_go/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type IReviewController interface {
	CreateReview(ctx *gin.Context)
	UpdateReview(ctx *gin.Context)
	GetReviewById(ctx *gin.Context)
	ListReviewsByProperty(ctx *gin.Context)
	DeleteReview(ctx *gin.Context)
}

type ReviewsController struct {
	ReviewsService iservice.IReviewService
}

// CreateReview implements IReviewController.
func (r *ReviewsController) CreateReview(ctx *gin.Context) {
	var in requestDTO.ReviewCreateModel

	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.BadResponse(ctx, response.CannotCreatePropertyCode)
		return
	}

	userID := auth.GetCurrentUserId(ctx)
	out, code, err := r.ReviewsService.CreateReview(ctx, in, userID)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

// DeleteReview implements IReviewController.
func (r *ReviewsController) DeleteReview(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), "/")
	code, err := r.ReviewsService.DeleteReview(ctx, id)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, nil)
}

// GetReviewById implements IReviewController.
func (r *ReviewsController) GetReviewById(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), "/")
	out, code, err := r.ReviewsService.GetReviewByID(ctx, id)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)

}

// ListReviewsByProperty implements IReviewController.
func (r *ReviewsController) ListReviewsByProperty(ctx *gin.Context) {
	var in database.ListReviewsByPropertyParams
	in.PropertyID = strings.Trim(ctx.Param("id"), "/")
	in.Limit, in.Offset = query.GetLimitAndOffsetFromQuery(ctx)

	out, code, err := r.ReviewsService.GetReviewByPropertyID(ctx, in)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

// UpdateReview implements IReviewController.
func (r *ReviewsController) UpdateReview(ctx *gin.Context) {
	var in requestDTO.ReviewUpdateModel

	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.BadResponse(ctx, response.CannotCreatePropertyCode)
		return
	}
	propertyID := strings.Trim(ctx.Param("id"), "/")

	out, code, err := r.ReviewsService.UpdateReview(ctx, in, propertyID)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

func NewReviewController(ReviewService iservice.IReviewService) IReviewController {
	return &ReviewsController{ReviewsService: ReviewService}
}
