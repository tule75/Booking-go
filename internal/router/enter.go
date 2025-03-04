package Router

import (
	"ecommerce_go/internal/router/manage"
	"ecommerce_go/internal/router/user"
)

type RouterGroup struct {
	UserGroupRouter  user.UserGroupRouter
	AdminGroupRouter manage.AdminGroupRouter
}

var RouterAppGroup = new(RouterGroup)
