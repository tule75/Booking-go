package iservice

import (
	"context"
	"ecommerce_go/internal/database"
	requestDTO "ecommerce_go/internal/models/request"
)

type IRoomService interface {
	CreateRoom(ctx context.Context, in requestDTO.RoomCreateModel) (string, int, error)
	UpdateRoom(ctx context.Context, in requestDTO.RoomUpdateModel, roomID string) (string, int, error)
	GetRoomByID(ctx context.Context, id string) (database.GetRoomByIDRow, int, error)
	GetRoomByPropertyID(ctx context.Context, in database.ListRoomsByPropertyParams) ([]database.ListRoomsByPropertyRow, int, error)
	DeleteRoom(ctx context.Context, id string) (int, error)
}

var (
	roomService IRoomService
)

func InitRoom(p IRoomService) {
	roomService = p
}

func GetRoom() IRoomService {
	if roomService == nil {
		panic("implement RoomService not found for interface IRoomService")
	}
	return roomService
}
