package models

type Message struct {
	Id      uint   `json:"id"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
