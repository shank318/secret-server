package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"secret-server/app/constants"
	"secret-server/app/utils"
	"testing"
)

func TestSetRequestPayload(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("GET", "/abc", nil)
	ctx.Request.RequestURI = "myuri"
	ctx.Request.Host = "myhost"
	ctx.Request.Header.Set(constants.X_RAZORPAY_TASK_ID, "TaskId")
	ctx.Set(constants.Mode, "Mode")
	handler := SetRequestPayload()
	handler(ctx)
	r, b := ctx.Get(constants.Request)
	assert.True(t, b)
	rm := r.(map[string]interface{})
	assert.Equal(t, "myuri", rm["uri"])
	assert.Equal(t, "myhost", rm["host"])
	assert.Equal(t, "TaskId", rm[constants.TaskID])
	assert.True(t, utils.IsValidUUID(rm[constants.RequestID].(string)))
}
