package models

type Base struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
