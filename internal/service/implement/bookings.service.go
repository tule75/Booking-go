package service

import (
	"context"
	"database/sql"
	"ecommerce_go/global"
	"ecommerce_go/internal/database"
	requestDTO "ecommerce_go/internal/models/request"
	iservice "ecommerce_go/internal/service/interface"
	"ecommerce_go/internal/utils/kafka"
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

type BookingService struct {
	sqlc *database.Queries
}

// CreateBooking implements iservice.IBookingService.
func (r *BookingService) CreateBooking(ctx context.Context, in requestDTO.BookingCreateModel, userID string) (string, int, error) {
	var booking = database.CreateBookingParams{
		ID:         uuid.New().String(),
		UserID:     userID,
		PropertyID: in.PropertyID,
		RoomID:     in.RoomID,
		CheckIn:    in.CheckIn,
		CheckOut:   in.CheckOut,
		Guests:     in.Guests,
		TotalPrice: in.TotalPrice,
		Status:     in.Status,
	}
	isLock := redis.CheckSoftLock(booking.RoomID.String, booking.UserID, booking.CheckIn, booking.CheckOut)

	if isLock {
		fmt.Println(constant.RoomIsLock)
		return "", response.BookingIsLocked, fmt.Errorf(constant.RoomIsLock)
	}

	kafka.SendInsertBooking(ctx, &booking, func(err error) {
		if err != nil {
			global.Logger.Error(constant.KafkaFailure, zap.Error(err))
		} else {
			global.Logger.Info(constant.KafkaSuccess)
		}
	})

	return booking.ID, response.SuccessResponseCode, nil
}

// DeleteBooking implements iservice.IBookingService.
func (r *BookingService) DeleteBooking(ctx context.Context, id string) (code int, err error) {
	err = r.sqlc.SoftDeleteBooking(ctx, id)

	if err != nil {
		global.Logger.Error(constant.DeleteBookingFailure, zap.Error(err))
		return response.CannotDeleteBookingCode, err
	}
	global.Logger.Info(constant.DeleteBookingSuccess, zap.Any("room id::", id))

	defer redis.DeleteCache(ctx, constant.PreRoomById, id)
	return response.SuccessResponseCode, nil
}

func (r *BookingService) CancelBooking(ctx context.Context, id string) (code int, err error) {
	err = r.sqlc.CancelBooking(ctx, id)

	if err != nil {
		global.Logger.Error(constant.CancelBookingFailure, zap.Error(err))
		return response.CannotDeleteBookingCode, err
	}
	global.Logger.Info(constant.CancelBookingSuccess, zap.Any("room id::", id))
	booking, _ := r.sqlc.GetBookingByID(ctx, id)

	kafka.SendUpdateAvailibility(ctx, &kafka.UpdateAvailabilityKafka{
		RoomID:      booking.RoomID.String,
		CheckIn:     booking.CheckIn,
		CheckOut:    booking.CheckOut,
		IsAvailable: true,
	}, func(err error) {
		if err != nil {
			global.Logger.Error(constant.KafkaFailure, zap.Error(err))
		} else {
			global.Logger.Info(constant.KafkaSuccess)
		}
	})

	defer redis.DeleteCache(ctx, constant.PreRoomById, id)
	return response.SuccessResponseCode, nil
}

// GetBookingByID implements iservice.IBookingService.
func (r *BookingService) GetBookingByID(ctx context.Context, id string) (out database.GetBookingByIDRow, code int, err error) {
	var booking database.GetBookingByIDRow
	// redis get data
	result, err := redis.GetSearchResultsFromRedis(ctx, constant.PreBookingById, id)
	if err == nil {
		// if error then countinue else return
		if err := json.Unmarshal([]byte(result), &booking); err == nil {
			return booking, response.SuccessResponseCode, nil
		}
	}

	booking, err = r.sqlc.GetBookingByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No property was found by ID:", id)
			return out, response.NotFoundResponseCode, nil
		}
		fmt.Printf("Error getting room:: %v\n", err)
		return out, response.CannotGetRoomByIDCode, err
	}

	go redis.CacheStore(ctx, constant.PreBookingById, id, booking, 5*time.Minute)
	return booking, response.SuccessResponseCode, nil
}

// GetBookingByUserID implements iservice.IBookingService.
func (r *BookingService) GetBookingByUserID(ctx context.Context, in database.ListBookingsByUserParams) (out []database.ListBookingsByUserRow, code int, err error) {
	// redis get data
	result, err := redis.GetSearchResultsFromRedis(ctx, constant.PreBookingByUserId, in.UserID)
	if err == nil {
		// if error then countinue else return
		if err := json.Unmarshal([]byte(result), &out); err == nil {
			return out, response.SuccessResponseCode, nil
		}
	}

	out, err = r.sqlc.ListBookingsByUser(ctx, in)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No property was found by ID:", in.UserID)
			return out, response.NotFoundResponseCode, nil
		}
		fmt.Printf("Error getting room:: %v\n", err)
		return out, response.CannotGetRoomByIDCode, err
	}

	go redis.CacheStore(ctx, constant.PreBookingByUserId, in.UserID, out, 5*time.Minute)
	return out, response.SuccessResponseCode, nil
}

// UpdateBooking implements iservice.IBookingService.
func (r *BookingService) UpdateBooking(ctx context.Context, in requestDTO.BookingUpdateModel, BookingID string) (string, int, error) {
	// var updateRoom = database.UpdateBookingParams{
	// 	ID:      BookingID,
	// 	CheckIn: ,
	// }

	// err := r.sqlc.UpdateBooking(ctx, updateRoom)
	// if err != nil {
	// 	global.Logger.Error("Error updating room", zap.Error(err))
	// 	return "", response.CannotCreatePropertyCode, err
	// }
	// global.Logger.Info("Update Room success::", zap.Any("value::", updateRoom))

	// defer redis.DeleteCache(ctx, constant.PreRoomByPropertiesId, in.PropertyID)
	// defer redis.DeleteCache(ctx, constant.PreRoomById, BookingID)

	// return updateRoom.ID, response.SuccessResponseCode, nil
	panic("UnImplemented")
}

func NewBookingService(sqlc *database.Queries) iservice.IBookingService {
	return &BookingService{sqlc: sqlc}
}
