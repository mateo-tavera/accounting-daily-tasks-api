package entity

import "time"

type Task struct {
	Id      int       `json:"id,omitempty"`
	User    string    `json:"user" validate:"required"`
	Summary string    `json:"summary" validate:"required"` // TODO: encrypt
	Date    time.Time `json:"date"`
	Status  string    `json:"status" validate:"required"`
}

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
