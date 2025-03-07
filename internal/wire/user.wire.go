//go:build wireinject

package wire

import (
	"ecommerce_go/internal/controller"
	"ecommerce_go/internal/repo"
	service "ecommerce_go/internal/service/implement"

	"github.com/google/wire"
)

func InitUserRouterHanlder() (controller.IUserController, error) {
	wire.Build(
		repo.NewUserRepo,
		repo.NewUserAuthRepo,
		service.NewUserLogin,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
