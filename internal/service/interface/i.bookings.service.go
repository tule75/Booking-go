package iservice

import (
	"context"
	"ecommerce_go/internal/database"
	requestDTO "ecommerce_go/internal/models/request"
)

type IBookingService interface {
	CreateBooking(ctx context.Context, in requestDTO.BookingCreateModel, userID string) (string, int, error)
	UpdateBooking(ctx context.Context, in requestDTO.BookingUpdateModel, BookingID string) (string, int, error)
	GetBookingByID(ctx context.Context, id string) (database.GetBookingByIDRow, int, error)
	GetBookingByUserID(ctx context.Context, in database.ListBookingsByUserParams) ([]database.ListBookingsByUserRow, int, error)
	DeleteBooking(ctx context.Context, id string) (int, error)
	CancelBooking(ctx context.Context, id string) (code int, err error)
}

var (
	BookingService IBookingService
)

func InitBooking(p IBookingService) {
	BookingService = p
}

func GetBooking() IBookingService {
	if BookingService == nil {
		panic("implement BookingService not found for interface IBookingService")
	}
	return BookingService
}
