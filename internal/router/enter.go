package Router

import (
	"ecommerce_go/internal/router/manage"
	"ecommerce_go/internal/router/properties"
	"ecommerce_go/internal/router/rooms"
	"ecommerce_go/internal/router/user"
)

type RouterGroup struct {
	UserGroupRouter       user.UserGroupRouter
	AdminGroupRouter      manage.AdminGroupRouter
	PropertiesGroupRouter properties.PropertyGroupRouter
	RoomGroupRouter       rooms.RoomGroupRouter
}

var RouterAppGroup = new(RouterGroup)
