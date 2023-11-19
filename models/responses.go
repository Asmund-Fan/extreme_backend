package models

import "calculator_backend/db"

type CalcHistoriesResponse struct {
	Data []db.CalcHistory `json:"data"`
}

type RateRecordResponse struct {
	Data db.RateRecord `json:"data"`
}

type UserResponse struct {
	Data db.User `json:"data"`
}
