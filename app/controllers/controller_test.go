package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"secret-server/app/constants"
	"secret-server/app/routing/middleware"
	"secret-server/app/routing/response"
	"secret-server/app/crerrors"
	"testing"
)

type databaseMock struct {
}

func (dbProvider *databaseMock) Ping(ctx *gin.Context) crerrors.IError {
	return nil
}

func TestAppController_Welcome(t *testing.T) {
	router := setupRouter()
	w := performRequest(router, "GET", "/", nil)
	resp := make(map[string]interface{})
	json.Unmarshal([]byte(w.Body.String()), &resp)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Welcome to Router :)", resp["Message"])
}

func TestStatusController_Ping(t *testing.T) {
	router := setupRouter()
	w := performRequest(router, "GET", "/ping", nil)
	assert.Equal(t, http.StatusOK, w.Code)
	resp := make(map[string]interface{})
	err := json.Unmarshal([]byte(w.Body.String()), &resp)
	if err != nil {
		assert.Fail(t, "Invalid response")
		return
	}
	assert.NotNil(t, resp["commit_id"])
	assert.NotNil(t, resp["container_id"])
}

func TestStatusController_Status(t *testing.T) {
	router := setupRouter()
	w := performRequest(router, "GET", "/status", nil)
	assert.Equal(t, http.StatusOK, w.Code)
}

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.SerializeResponse, middleware.Recovery)
	rootGroup := router.Group("/")
	rootGroup.GET("", RequestResponseHandler(App.Welcome))
	rootGroup.GET("status", RequestResponseHandler(App.Status))
	rootGroup.GET("ping", RequestResponseHandler(App.Ping))
	return router
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func RequestResponseHandler(handler handlerFunc) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		res := handler(ctx)
		ctx.Set(constants.Response, res)
	}
}

type handlerFunc func(ctx *gin.Context) response.Response
