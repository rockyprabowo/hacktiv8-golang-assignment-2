package responses

// WithDataList is the list of item T response payload.
type WithDataList[T any] struct {
	Count   int    `json:"count"`
	Data    []T    `json:"data"`
	Message string `json:"message"`
}

// WithData is the single item T response payload.
type WithData[T any] struct {
	Data   T      `json:"data"`
	Status string `json:"message"`
}
