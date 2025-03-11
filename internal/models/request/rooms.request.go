package requestDTO

type RoomCreateModel struct {
	PropertyID  string `json:"property_id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	MaxGuests   int    `json:"max_guests"`
	IsAvailable bool   `json:"is_available"`
}

type RoomUpdateModel struct {
	Name        string `json:"name"`
	Price       string `json:"price"`
	MaxGuests   int    `json:"max_guests"`
	IsAvailable bool   `json:"is_available"`
	PropertyID  string `json:"property_id"`
	ID          string `json:"id"`
}
