package initialize

import (
	"ecommerce_go/internal/database"
	service "ecommerce_go/internal/service/implement"
	iservice "ecommerce_go/internal/service/interface"
)

func InitServices(queries *database.Queries) {
	// User Service Interface
	iservice.InitUserLogin(service.NewUserLogin(queries))
	iservice.InitProperty(service.NewPropertiesService(queries))
	iservice.InitRoom(service.NewRoomService(queries))
	iservice.InitReview(service.NewReviewService(queries))
	iservice.InitBooking(service.NewBookingService(queries))
}
