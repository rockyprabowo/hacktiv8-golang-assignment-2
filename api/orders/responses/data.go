package order_responses

import . "rocky.my.id/git/h8-assignment-2/models"

// DataList is the list of order response payload.
type DataList struct {
	Count   int     `json:"count"`
	Data    []Order `json:"data"`
	Message string  `json:"message"`
}

// Data is the single order response payload.
type Data struct {
	Data    Order  `json:"data"`
	Message string `json:"message"`
}
