package order_responses

// RowsAffected is a response payload for any operations with affected row count
// from the database.
type RowsAffected struct {
	Count   int    `json:"count"`
	Message string `json:"message"`
}
