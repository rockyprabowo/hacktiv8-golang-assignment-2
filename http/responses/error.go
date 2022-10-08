package responses

// Error
// @Description Represents the default error response payload.
type Error struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"something is wrong"`
}
