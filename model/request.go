package model

type NewUrlRequest struct {
	ActualUrl string `json:"actual_url"`
}

type UpdateStatusRequest struct {
	ID     int  `json:"id"`
	Status bool `json:"status"`
}
