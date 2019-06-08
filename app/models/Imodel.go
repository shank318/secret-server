package models

import "github.com/go-ozzo/ozzo-validation"

type IModel interface {
	validation.Validatable
}
