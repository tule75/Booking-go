package requestDTO

type PropertyCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Location    string `json:"location" binding:"required"`
	Price       string `json:"price" binding:"required"`
	Amenities   string `json:"amenities" binding:"required"`
}

type PropertyUpdateRequest struct {
	OwnerID     string `json:"owner_id", binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Location    string `json:"location" binding:"required"`
	Price       string `json:"price" binding:"required"`
	Amenities   string `json:"amenities" binding:"required"`
}
