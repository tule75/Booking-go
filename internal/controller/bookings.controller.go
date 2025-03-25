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

type IBookingsController interface {
	CreateBooking(ctx *gin.Context)
	UpdateBooking(ctx *gin.Context)
	GetBookingById(ctx *gin.Context)
	ListBookingsByUser(ctx *gin.Context)
	CancelBooking(ctx *gin.Context)
	DeleteBooking(ctx *gin.Context)
}

type BookingsController struct {
	BookingService iservice.IBookingService
}

// CreateBooking implements IBookingController.
func (r *BookingsController) CreateBooking(ctx *gin.Context) {
	var in requestDTO.BookingCreateModel

	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.BadResponse(ctx, response.CannotCreatePropertyCode)
		return
	}

	userID := auth.GetCurrentUserId(ctx)

	out, code, err := r.BookingService.CreateBooking(ctx, in, userID)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

// DeleteBooking implements IBookingController.
func (r *BookingsController) DeleteBooking(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), "/")
	code, err := r.BookingService.DeleteBooking(ctx, id)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, nil)
}

// CancelBooking implements IBookingController.
func (r *BookingsController) CancelBooking(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), "/")
	code, err := r.BookingService.CancelBooking(ctx, id)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, nil)
}

// GetBookingById implements IBookingController.
func (r *BookingsController) GetBookingById(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), "/")
	out, code, err := r.BookingService.GetBookingByID(ctx, id)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)

}

// ListBookingsByUser implements IBookingController.
func (r *BookingsController) ListBookingsByUser(ctx *gin.Context) {
	var in database.ListBookingsByUserParams
	in.UserID = strings.Trim(ctx.Param("id"), "/")
	in.Limit, in.Offset = query.GetLimitAndOffsetFromQuery(ctx)

	out, code, err := r.BookingService.GetBookingByUserID(ctx, in)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

// UpdateBooking implements IBookingController.
func (r *BookingsController) UpdateBooking(ctx *gin.Context) {
	var in requestDTO.BookingUpdateModel

	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.BadResponse(ctx, response.CannotCreatePropertyCode)
		return
	}
	bookingID := strings.Trim(ctx.Param("id"), "/")

	out, code, err := r.BookingService.UpdateBooking(ctx, in, bookingID)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

func NewBookingController(BookingService iservice.IBookingService) IBookingsController {
	return &BookingsController{BookingService: BookingService}
}
