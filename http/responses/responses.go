package responses

// WithDataList
// @Description Represents the list of item T response payload.
type WithDataList[T any] struct {
	Count   int    `json:"count" example:"1"`
	Data    []T    `json:"data"`
	Message string `json:"message" example:"success"`
}

// WithSingleData
// @Description Represents the single item T response payload.
type WithSingleData[T any] struct {
	Data   T      `json:"data"`
	Status string `json:"message" example:"success"`
}
