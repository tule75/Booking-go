package reviews

import (
	"ecommerce_go/internal/controller"
	"ecommerce_go/internal/middleware"
	iservice "ecommerce_go/internal/service/interface"

	"github.com/gin-gonic/gin"
)

type ReviewsRouter struct {
}

func (r *ReviewsRouter) InitPropertiesRouter(Router *gin.RouterGroup) {
	rc := controller.NewReviewController(iservice.GetReview())
	userPublicRouter := Router.Group("/reviews")
	{
		userPublicRouter.GET("/property/:id", rc.ListReviewsByProperty)
		userPublicRouter.GET("/:id", rc.GetReviewById)
	}

	userPrivateRouter := Router.Group("/users")
	userPrivateRouter.Use(middleware.AuthenticationMiddleware())
	userPrivateRouter.Use(middleware.Authorization([]string{"ADMIN", "CUSTOMER"}))
	{
		userPrivateRouter.POST("/", rc.CreateReview)
		userPrivateRouter.PUT("/", rc.UpdateReview)
		userPrivateRouter.DELETE("/", rc.DeleteReview)
	}
}
