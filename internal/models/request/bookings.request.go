package requestDTO

import (
	"database/sql"
	"ecommerce_go/internal/database"
	"time"
)

type BookingCreateModel struct {
	PropertyID string                      `json:"property_id"`
	RoomID     sql.NullString              `json:"room_id"`
	CheckIn    time.Time                   `json:"check_in"`
	CheckOut   time.Time                   `json:"check_out"`
	Guests     int32                       `json:"guests"`
	TotalPrice string                      `json:"total_price"`
	Status     database.NullBookingsStatus `json:"status"`
}

type BookingUpdateModel struct {
	CheckIn    time.Time                   `json:"check_in"`
	CheckOut   time.Time                   `json:"check_out"`
	Guests     int32                       `json:"guests"`
	TotalPrice string                      `json:"total_price"`
	Status     database.NullBookingsStatus `json:"status"`
}
