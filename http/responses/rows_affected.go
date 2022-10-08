package responses

// WithRowsAffected
// @Description Represents a response payload for any operations
// @Description with affected row count from the database.
type WithRowsAffected struct {
	Count   int    `json:"count" example:"13"`
	Message string `json:"message" example:"success"`
}
