package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"time"
)

type Secret struct {
	ID             string
	SecretText     string
	CreatedAt      time.Time
	ExpiresAt      time.Time
	MaxViews       int
	RemViews int
}

type SecretResponse struct {
	Hash           string    `xml:"hash,attr" json:"hash"`
	SecretText     string    `xml:"secretText,attr" json:"secretText"`
	CreatedAt      time.Time `xml:"createdAt,attr" json:"createdAt"`
	ExpiresAt      time.Time `xml:"expiresAt,attr" json:"expiresAt"`
	RemainingViews int `xml:"remainingViews,attr" json:"remainingViews"`
}

type SecretRequest struct {
	SecretText       string
	ExpireAfterViews int
	ExpireAfter      int
	IModel
}

func (s SecretRequest) Validate() error {
	return validation.ValidateStruct(
		&s,
		validation.Field(&s.SecretText, validation.Required),
		validation.Field(&s.ExpireAfterViews, validation.Required),
		validation.Field(&s.ExpireAfterViews, validation.Required),
	)
}
