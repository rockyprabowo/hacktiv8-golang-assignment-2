package responses

// WithRowsAffected is a response payload for any operations with affected row count
// from the database.
type WithRowsAffected struct {
	Count   int    `json:"count"`
	Message string `json:"message"`
}
