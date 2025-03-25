package reviews

import (
	"ecommerce_go/internal/controller"
	"ecommerce_go/internal/middleware"
	iservice "ecommerce_go/internal/service/interface"
	constant "ecommerce_go/pkg"

	"github.com/gin-gonic/gin"
)

type ReviewsRouter struct {
}

func (r *ReviewsRouter) InitReviewsRouter(Router *gin.RouterGroup) {
	rc := controller.NewReviewController(iservice.GetReview())
	reviewPublicRouter := Router.Group("/reviews")
	{
		reviewPublicRouter.GET("/property/:id", rc.ListReviewsByProperty)
		reviewPublicRouter.GET("/:id", rc.GetReviewById)
	}

	reviewPrivateRouter := Router.Group("/reviews")
	reviewPrivateRouter.Use(middleware.AuthenticationMiddleware())
	reviewPrivateRouter.Use(middleware.Authorization([]string{constant.RoleAdmin, constant.RoleCustomer}))
	{
		reviewPrivateRouter.POST("/", rc.CreateReview)
		reviewPrivateRouter.PUT("/:id", rc.UpdateReview)
		reviewPrivateRouter.DELETE("/", rc.DeleteReview)
	}
}
