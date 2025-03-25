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
// Create implements IRoomController.
// Create godoc
// @Summary      Create Room
// @Description  Create Room
// @Tags         Room management
// @Accept       json
// @Produce      json
// @Param        payload body requestDTO.RoomCreateModel true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /rooms [POST]
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
// Delete implements IRoomController.
// Delete godoc
// @Summary      Delete a room by ID
// @Description  Deletes a room from the system using its unique ID.
// @Tags         Room management
// @Accept       json
// @Produce      json
// @Param        id path string true "Room ID"
// @Success      200 {object} response.ResponseData "Room deleted successfully"
// @Router       /rooms/{id} [delete]
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
// Get implements IRoomController.
// Get godoc
// @Summary      Get a room by ID
// @Description  Get a room from the system using its unique ID.
// @Tags         Room management
// @Accept       json
// @Produce      json
// @Param        id path string true "Room ID"
// @Success      200 {object} response.ResponseData
// @Router       /rooms/{id} [GET]
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
// Get implements IRoomController.
// Get godoc
// @Summary      Get rooms by property ID
// @Description  Get rooms from the system using property unique ID.
// @Tags         Room management
// @Accept       json
// @Produce      json
// @Param        id path string true "Property ID"
// @Param        limit query int false "Limit the number of rooms affected (default: 20)"
// @Param        offset query int false "Offset for starting position (default: 0)"
// @Success      200 {object} response.ResponseData
// @Router       /rooms/property/{id} [GET]
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
// Update implements IRoomController.
// Update godoc
// @Summary      Update room by ID
// @Description  Update rooms from the system using its unique ID.
// @Tags         Room management
// @Accept       json
// @Produce      json
// @Param        id path string true "Property ID"
// @Param        payload body requestDTO.RoomUpdateModel true "payload"
// @Success      200 {object} response.ResponseData
// @Router       /rooms/{id} [PUT]
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
