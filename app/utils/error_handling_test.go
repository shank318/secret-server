package utils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"secret-server/app/crerrors"
	"testing"
)

func TestGetError(t *testing.T) {
	e := GetError("Error 1")
	assert.Equal(t, "Error 1", e.Error())

	err := errors.New("Error 2")
	assert.Equal(t, err, GetError(err))

	assert.Equal(t, crerrors.InternalServerErrorCode, GetError(nil).Error())

	rzp := crerrors.NewCrError(nil, crerrors.InternalServerErrorCode, err)
	assert.Equal(t, rzp, GetError(rzp))
}
