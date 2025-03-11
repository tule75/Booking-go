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

type IPropertiesController interface {
	CreateProperty(ctx *gin.Context)
	UpdateProperty(ctx *gin.Context)
	GetPropertyByID(ctx *gin.Context)
	GetPropertiesByOwnerID(ctx *gin.Context)
	SearchProperties(ctx *gin.Context)
	DeleteProperty(ctx *gin.Context)
}

type PropertiesController struct {
	PropertiesService iservice.IPropertiesService
}

// CreateProperty implements IPropertiesController.
// Register implements IPropertiesController.
// Register godoc
// @Summary      Create a new property
// @Description  Create a new property
// @Tags         properties management
// @Accept       json
// @Produce      json
// @Param        payload body requestDTO.PropertyCreateRequest true "param"
// @Success      200  {object}  response.ResponseData
// @Router       /properties [POST]
func (p *PropertiesController) CreateProperty(ctx *gin.Context) {
	var in requestDTO.PropertyCreateRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.BadResponse(ctx, response.CannotCreatePropertyCode)
		return
	}
	userID := auth.GetCurrentUserId(ctx)

	out, code, err := p.PropertiesService.CreateProperty(ctx, in, userID)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

// GetPropertiesByOwnerID implements IPropertiesController.
// Register implements IPropertiesController.
// Register godoc
// @Summary      GetPropertiesByOwnerID
// @Description  GetPropertiesByOwnerID
// @Tags         properties management
// @Accept       json
// @Produce      json
// @Param        payload body requestDTO.PropertyCreateRequest true "payload"
// @Success      200  {object}  response.ResponseData
// @Router       /properties/owner/:id [GET]
func (p *PropertiesController) GetPropertiesByOwnerID(ctx *gin.Context) {
	var in database.ListPropertiesByOwnerParams

	in.OwnerID = strings.Trim(ctx.Param("owner_id"), "/")
	in.Limit, in.Offset = query.GetLimitAndOffsetFromQuery(ctx)

	if in.OwnerID == "" || strings.ToLower(in.OwnerID) == "current_user" {
		userID := auth.GetCurrentUserId(ctx)
		in.OwnerID = userID
	}

	out, code, err := p.PropertiesService.GetPropertiesByOwner(ctx, in)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

// GetPropertyByID implements IPropertiesController.
func (p *PropertiesController) GetPropertyByID(ctx *gin.Context) {
	var id = strings.Trim(ctx.Param("id"), "/")

	out, code, err := p.PropertiesService.GetProperty(ctx, id)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
	}

	response.SuccessResponse(ctx, code, out)
}

// SearchProperties implements IPropertiesController.
func (p *PropertiesController) SearchProperties(ctx *gin.Context) {
	var in database.SearchPropertiesParams
	in.Limit, in.Offset = query.GetLimitAndOffsetFromQuery(ctx)
	in.FromPrice = ctx.Query("fromprice")
	in.ToPrice = ctx.Query("toprice")

	out, code, err := p.PropertiesService.SearchProperties(ctx, in)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
	}

	response.SuccessResponse(ctx, code, out)
}

// UpdateProperty implements IPropertiesController.
func (p *PropertiesController) UpdateProperty(ctx *gin.Context) {
	var in requestDTO.PropertyUpdateRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		response.BadResponse(ctx, response.CannotCreatePropertyCode)
		return
	}
	userID := auth.GetCurrentUserId(ctx)
	in.OwnerID = userID

	out, code, err := p.PropertiesService.UpdateProperty(ctx, in)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, out)
}

// DeletePropertie implements IPropertyController.
func (p *PropertiesController) DeleteProperty(ctx *gin.Context) {
	id := strings.Trim(ctx.Param("id"), "/")
	code, err := p.PropertiesService.DeleteProperty(ctx, id)

	if err != nil {
		response.ErrorResponse(ctx, code, err.Error())
		return
	}

	response.SuccessResponse(ctx, code, nil)
}

func NewPropertiesController(propertiesService iservice.IPropertiesService) IPropertiesController {
	return &PropertiesController{PropertiesService: propertiesService}
}
