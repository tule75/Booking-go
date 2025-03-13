package controller

import (
	"ecommerce_go/internal/database"
	requestDTO "ecommerce_go/internal/models/request"
	iservice "ecommerce_go/internal/service/interface"
	"ecommerce_go/internal/utils/query"
	"ecommerce_go/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type IRoomController interface {
	CreateRoom(ctx *gin.Context)
	UpdateRoom(ctx *gin.Context)
	GetRoomById(ctx *gin.Context)
	ListRoomsByProperty(ctx *gin.Context)
	DeleteRoom(ctx *gin.Context)
}

type RoomsController struct {
	RoomsService iservice.IRoomService
}

// CreateRoom implements IRoomController.
func (r *RoomsController) CreateRoom(ctx *gin.Context) {
	var in requestDTO.RoomCreateModel

	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.BadResponse(ctx, response.CannotCreatePropertyCode)
		return
	}

	out, code, err := r.RoomsService.CreateRoom(ctx, in)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

// DeleteRoom implements IRoomController.
func (r *RoomsController) DeleteRoom(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), "/")
	code, err := r.RoomsService.DeleteRoom(ctx, id)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, nil)
}

// GetRoomById implements IRoomController.
func (r *RoomsController) GetRoomById(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), "/")
	out, code, err := r.RoomsService.GetRoomByID(ctx, id)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)

}

// ListRoomsByProperty implements IRoomController.
func (r *RoomsController) ListRoomsByProperty(ctx *gin.Context) {
	var in database.ListRoomsByPropertyParams
	in.PropertyID = strings.Trim(ctx.Param("id"), "/")
	in.Limit, in.Offset = query.GetLimitAndOffsetFromQuery(ctx)

	out, code, err := r.RoomsService.GetRoomByPropertyID(ctx, in)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

// UpdateRoom implements IRoomController.
func (r *RoomsController) UpdateRoom(ctx *gin.Context) {
	var in requestDTO.RoomUpdateModel

	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.BadResponse(ctx, response.CannotCreatePropertyCode)
		return
	}
	propertyID := strings.Trim(ctx.Param("id"), "/")

	out, code, err := r.RoomsService.UpdateRoom(ctx, in, propertyID)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

func NewRoomController(roomService iservice.IRoomService) IRoomController {
	return &RoomsController{RoomsService: roomService}
}
