package utils

import (
	"errors"
	"secret-server/app/crerrors"
)

// GetError Gets error interface given the interface of error
// Extract the error data from the err type
// If type couldn't be determined then create a default error
func GetError(err interface{}) error {
	var er error

	switch err.(type) {
	case error:
		er = err.(error)

	case crerrors.IError:
		er = err.(crerrors.IError)

	case string:
		er = errors.New(err.(string))

	default:
		er = errors.New(crerrors.InternalServerErrorCode)
	}

	return er
}
