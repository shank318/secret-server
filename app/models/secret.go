package models

import "time"

type Secret struct {
	ID             string
	SecretText     string
	CreatedAt      time.Time
	ExpiresAt      time.Time
	MaxViews       int
	RemainingViews time.Time
}

type SecretResponse struct {
	Hash           string
	SecretText     string
	CreatedAt      time.Time
	ExpiresAt      time.Time
	RemainingViews time.Time
}
