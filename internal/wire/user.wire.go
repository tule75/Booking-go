//go:build wireinject

package wire

import (
	"ecommerce_go/internal/controller"
	"ecommerce_go/internal/repo"
	"ecommerce_go/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHanlder() (controller.IUserController, error) {
	wire.Build(
		repo.NewUserRepo,
		repo.NewUserAuthRepo,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
