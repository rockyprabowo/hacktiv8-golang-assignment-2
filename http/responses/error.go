package responses

// Error is the default error response payload.
type Error struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
