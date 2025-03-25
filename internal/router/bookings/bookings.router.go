package bookings

import (
	"ecommerce_go/internal/controller"
	"ecommerce_go/internal/middleware"
	iservice "ecommerce_go/internal/service/interface"
	constant "ecommerce_go/pkg"

	"github.com/gin-gonic/gin"
)

type BookingsRouter struct {
}

func (r *BookingsRouter) InitBookingRouter(Router *gin.RouterGroup) {
	rc := controller.NewBookingController(iservice.GetBooking())

	bookingPrivateRouter := Router.Group("/bookings")
	bookingPrivateRouter.Use(middleware.AuthenticationMiddleware())
	bookingPrivateRouter.Use(middleware.Authorization([]string{constant.RoleAdmin, constant.RoleCustomer}))
	{
		bookingPrivateRouter.POST("/", rc.CreateBooking)
		bookingPrivateRouter.PUT("/:id", rc.UpdateBooking)
		bookingPrivateRouter.DELETE("/", rc.DeleteBooking)
		bookingPrivateRouter.PUT("/cancel/:id", rc.CancelBooking)
		bookingPrivateRouter.GET("/:id", rc.GetBookingById)
		bookingPrivateRouter.GET("/get_by_uid/:id", rc.ListBookingsByUser)
	}
}
