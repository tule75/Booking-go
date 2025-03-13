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

type RoomService struct {
	sqlc *database.Queries
}

// CreateRoom implements iservice.IRoomService.
func (r *RoomService) CreateRoom(ctx context.Context, in requestDTO.RoomCreateModel) (string, int, error) {
	var room = database.CreateRoomParams{
		ID:          uuid.New().String(),
		PropertyID:  in.PropertyID,
		Name:        sql.NullString{String: in.Name, Valid: true},
		Price:       in.Price,
		MaxGuests:   int32(in.MaxGuests),
		IsAvailable: sql.NullBool{Bool: in.IsAvailable, Valid: true},
	}

	result, err := r.sqlc.CreateRoom(ctx, room)
	if err != nil {
		global.Logger.Error("Error creating room", zap.Error(err))
		return "", response.CannotCreateRoomCode, err
	}
	global.Logger.Info("new Room::", zap.Any("result::", result), zap.Any("value::", room))
	redis.DeleteCache(ctx, constant.PreRoomByPropertiesId, room.PropertyID)

	return room.ID, response.SuccessResponseCode, nil
}

// DeleteRoom implements iservice.IRoomService.
func (r *RoomService) DeleteRoom(ctx context.Context, id string) (code int, err error) {
	err = r.sqlc.SoftDeleteRoom(ctx, id)

	if err != nil {
		global.Logger.Error("Delete Room failed", zap.Error(err))
		return response.CannotDeleteRoomCode, err
	}
	global.Logger.Info("Delete Room success::", zap.Any("room id::", id))

	defer redis.DeleteCache(ctx, constant.PreRoomById, id)
	return response.SuccessResponseCode, nil
}

// GetRoomByID implements iservice.IRoomService.
func (r *RoomService) GetRoomByID(ctx context.Context, id string) (out database.GetRoomByIDRow, code int, err error) {
	var room database.GetRoomByIDRow
	// redis get data
	result, err := redis.GetSearchResultsFromRedis(ctx, constant.PreRoomById, id)
	if err == nil {
		// if error then countinue else return
		if err := json.Unmarshal([]byte(result), &room); err == nil {
			return room, response.SuccessResponseCode, nil
		}
	}

	room, err = r.sqlc.GetRoomByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No property was found by ID:", id)
			return out, response.NotFoundResponseCode, nil
		}
		fmt.Printf("Error getting room:: %v\n", err)
		return out, response.CannotGetRoomByIDCode, err
	}

	go redis.CacheStore(ctx, constant.PreRoomById, id, room, 5*time.Minute)
	return room, response.SuccessResponseCode, nil
}

// GetRoomByPropertyID implements iservice.IRoomService.
func (r *RoomService) GetRoomByPropertyID(ctx context.Context, in database.ListRoomsByPropertyParams) (out []database.ListRoomsByPropertyRow, code int, err error) {
	var rooms []database.ListRoomsByPropertyRow
	query := fmt.Sprintf("id:%s-limit:%v-offset:%v", in.PropertyID, in.Limit, in.Offset)
	result, err := redis.GetSearchResultsFromRedis(ctx, constant.PreRoomByPropertiesId, query)

	if err == nil {
		if err := json.Unmarshal([]byte(result), &rooms); err == nil {
			return rooms, response.SuccessResponseCode, nil
		}
	}

	rooms, err = r.sqlc.ListRoomsByProperty(ctx, database.ListRoomsByPropertyParams{})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("No property was found by property ID:", in.PropertyID)
			return out, response.NotFoundResponseCode, nil
		}
		fmt.Printf("Error listing rooms: %v", err)
		return out, response.CannotListRoomByPropertyCode, err
	}

	go redis.CacheStore(ctx, constant.PreRoomByPropertiesId, query, rooms, 5*time.Minute)
	return rooms, response.SuccessResponseCode, err
}

// UpdateRoom implements iservice.IRoomService.
func (r *RoomService) UpdateRoom(ctx context.Context, in requestDTO.RoomUpdateModel, roomID string) (string, int, error) {
	var updateRoom = database.UpdateRoomParams{
		ID:          roomID,
		Name:        sql.NullString{String: in.Name, Valid: true},
		Price:       in.Price,
		MaxGuests:   int32(in.MaxGuests),
		IsAvailable: sql.NullBool{Bool: in.IsAvailable, Valid: true},
	}

	err := r.sqlc.UpdateRoom(ctx, updateRoom)
	if err != nil {
		global.Logger.Error("Error updating room", zap.Error(err))
		return "", response.CannotCreatePropertyCode, err
	}
	global.Logger.Info("Update Room success::", zap.Any("value::", updateRoom))

	defer redis.DeleteCache(ctx, constant.PreRoomByPropertiesId, in.PropertyID)
	defer redis.DeleteCache(ctx, constant.PreRoomById, roomID)

	return updateRoom.ID, response.SuccessResponseCode, nil
}

func NewRoomService(sqlc *database.Queries) iservice.IRoomService {
	return &RoomService{sqlc: sqlc}
}
