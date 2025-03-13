package requestDTO

type ReviewCreateModel struct {
	UserID     string `json:"user_id"`
	PropertyID string `json:"property_id"`
	Rating     int32  `json:"rating"`
	Comment    string `json:"comment"`
}

type ReviewUpdateModel struct {
	PropertyID string `json:"property_id"`
	Rating     int    `json:"rating"`
	Comment    string `json:"comment"`
}
