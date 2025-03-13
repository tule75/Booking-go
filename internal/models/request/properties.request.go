package requestDTO

import (
	"database/sql"
	"encoding/json"
)

type PropertyCreateRequest struct {
	Name        string         `json:"name", binding:"required"`
	Description sql.NullString `json:"description"`
	Location    string         `json:"location", binding:"required"`
	Price       string         `json:"price", binding:"required"`
	Amenities   json.RawMessage
}

type PropertyUpdateRequest struct {
	Name        string          `json:"name", binding:"required"`
	OwnerID     string          `json:"owner_id", binding:"required"`
	Description sql.NullString  `json:"description", binding:"required"`
	Location    string          `json:"location", binding:"required"`
	Price       string          `json:"price", binding:"required"`
	Amenities   json.RawMessage `json:"menities", binding:"required"`
}
