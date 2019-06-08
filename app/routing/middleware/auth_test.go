package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"secret-server/app/config"
	"secret-server/app/constants"
	"testing"
)

func TestSecureCompare_WithDefinedInputs(t *testing.T) {
	assert.True(t, secureCompare("1234567890", "1234567890"))
	assert.False(t, secureCompare("123456789", "1234567890"))
	assert.False(t, secureCompare("12345678900", "1234567890"))
	assert.False(t, secureCompare("1234567891", "1234567890"))
}

func TestBasicAuth(t *testing.T) {
	conf := config.AppConfig{}
	conf.AuthUser.API.UserName = "api_ABC"
	conf.AuthUser.API.Password = "XYZ"
	oldAppConfig := appConfig
	defer func() { appConfig = oldAppConfig }()
	appConfig = func() config.AppConfig {
		return conf
	}
	var routeApps = []string{constants.APIAuthUserFieldName}
	handler := BasicAuth(routeApps)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("GET", "/", nil)
	ctx.Request.SetBasicAuth("api_ABC", "XYZ")
	handler(ctx)
	assert.False(t, ctx.IsAborted())
}
func TestBasicAuth_WrongCredentials(t *testing.T) {
	conf := config.AppConfig{}
	conf.AuthUser.API.UserName = "api_ABC"
	conf.AuthUser.API.Password = "XYZ"
	oldAppConfig := appConfig
	defer func() { appConfig = oldAppConfig }()
	appConfig = func() config.AppConfig {
		return conf
	}
	var routeApps = []string{constants.APIAuthUserFieldName}
	handler := BasicAuth(routeApps)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("GET", "/vhh", nil)
	ctx.Request.SetBasicAuth("api_AB", "XYZ")
	handler(ctx)
	SerializeResponse(ctx)
	assert.True(t, ctx.IsAborted())
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"code\":\"UNAUTHORIZED\",\"message\":\"Unauthorized: Invalid Username or Password\",\"sub_code\":\"\"}", w.Body.String())

}

func TestBasicAuth_InvalidApp(t *testing.T) {
	var routeApps = []string{constants.APIAuthUserFieldName}
	handler := BasicAuth(routeApps)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("GET", "/", nil)
	ctx.Request.SetBasicAuth("ABC", "XYZ")
	handler(ctx)
	SerializeResponse(ctx)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"code\":\"UNAUTHORIZED\",\"message\":\"Unauthorized: Invalid Username or Password\",\"sub_code\":\"\"}", w.Body.String())
}

func TestBasicAuth_WrongApp(t *testing.T) {
	var routeApps = []string{constants.APIAuthUserFieldName}
	handler := BasicAuth(routeApps)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("GET", "/", nil)
	ctx.Request.SetBasicAuth("PQR_ABC", "XYZ")
	handler(ctx)
	SerializeResponse(ctx)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"code\":\"UNAUTHORIZED\",\"message\":\"Unauthorized: Invalid Username or Password\",\"sub_code\":\"\"}", w.Body.String())
}

func TestBasicAuth_InvalidHeader(t *testing.T) {
	var routeApps = []string{constants.APIAuthUserFieldName}
	handler := BasicAuth(routeApps)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request, _ = http.NewRequest("GET", "/", nil)
	ctx.Request.Header.Set("Authorization", "ABCD")
	handler(ctx)
	SerializeResponse(ctx)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Equal(t, "{\"code\":\"UNAUTHORIZED\",\"message\":\"Unauthorized: Invalid Authorization Header: ABCD\",\"sub_code\":\"\"}", w.Body.String())
}
