package crerrors

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewRzpError(t *testing.T) {

	e := errors.New("Some error happened")
	err := NewCrError(&gin.Context{}, BadRequestErrorCode, e)

	code := err.ErrorCode()
	message := err.Message()
	er := err.Error()
	statusCode := err.StatusCode()
	ie := err.GetInternalError()
	td := err.GetTraceData()
	idetails := err.InternalErrorDetails()
	pc := err.PublicErrorCode()
	pm := err.PublicMessage()
	pdetails := err.PublicErrorDetails()
	subCode := err.SubCode()

	assert.Equal(t, "", subCode)

	assert.Equal(t, BadRequestErrorCode, code)
	assert.Equal(t, "Some error happened", message)
	assert.Equal(t, "Some error happened", er)
	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Equal(t, e, ie)
	assert.Equal(t, idetails, td[InternalErrorDetails])
	assert.Equal(t, pdetails, td[PublicErrorDetails])
	assert.Equal(t, "Some error happened", idetails[Message])
	assert.Equal(t, BadRequestErrorCode, idetails[Code])
	assert.Equal(t, BadRequestErrorCode, pc)
	assert.Equal(t, BadRequestErrorMessage, pm)
	assert.Equal(t, BadRequestErrorCode, pdetails[Code])
	assert.Equal(t, "invalid request sent", pdetails[Message])

	data := make(map[string]interface{})
	data["A"] = "B"

	err = err.WithFields(data)
	err = err.WithField("X", "Y")
	td = err.GetTraceData()
	assert.Equal(t, "Y", td["X"])
	assert.Equal(t, "B", td["A"])

	err = err.WithSubCode("CODE")
	assert.Equal(t, "CODE", err.SubCode())
	ds := make([]map[string]interface{}, 1)
	ds = append(ds, data)
	err.SetErrorData(ds)

	ds1 := err.ErrorData()
	assert.Equal(t, ds, ds1)

	assert.True(t, err.IsSubCode("CODE"))

	err.WithInternalMessage("My Message")
	td = err.GetTraceData()
	assert.Equal(t, "My Message", td["message"])
}
