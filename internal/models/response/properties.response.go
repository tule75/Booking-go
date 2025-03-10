package responseDTO

import (
	"database/sql"
	"encoding/json"
)

type PropertyResponse struct {
	PropertyID  string
	OwnerID     string
	Name        string
	Description sql.NullString
	Location    string
	Price       string
	Amenities   json.RawMessage
	CreatedAt   sql.NullTime
}
