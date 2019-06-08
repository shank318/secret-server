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
	Hash           string    `xml:"hash,attr" json:"hash"`
	SecretText     string    `xml:"secretText,attr" json:"secretText"`
	CreatedAt      time.Time `xml:"createdAt,attr" json:"createdAt"`
	ExpiresAt      time.Time `xml:"expiresAt,attr" json:"expiresAt"`
	RemainingViews time.Time `xml:"remainingViews,attr" json:"remainingViews"`
}
