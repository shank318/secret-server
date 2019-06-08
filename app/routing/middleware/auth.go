package middleware

import (
	"bytes"
	"crypto/subtle"
	"github.com/pkg/errors"
	"io"
	"secret-server/app/config"
	"secret-server/app/constants"
	"secret-server/app/routing/response"
	"secret-server/app/utils"
	"secret-server/app/crerrors"
	"strings"

	"github.com/gin-gonic/gin"
)

var appConfig = config.GetConfig

func respondWithError(code int, message string, ctx *gin.Context) {
	ierr := crerrors.NewCrError(ctx, crerrors.CodeUnauthorized, errors.New(message))
	responseStruct := response.
		NewResponse(ctx).
		SetError(ierr).
		SetStatusCode(code)
	ctx.Set(constants.Response, responseStruct)
	ctx.Abort()
}

func getHeader(key string, c *gin.Context) string {
	if values, _ := c.Request.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}

func readBody(reader io.Reader) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	return buf.String()
}

//TokenAuthMiddleware ...
func BasicAuth(routeApp []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestUsername, requestPassword, ok := c.Request.BasicAuth()

		if !ok {
			respondWithError(401, "Unauthorized: Invalid Authorization Header: "+c.Request.Header.Get("Authorization"), c)
			return
		}

		appName, status := isAuthorized(routeApp, requestUsername)

		if status == false {
			respondWithError(401, "Unauthorized: Invalid Username or Password", c)

			return
		}

		appUserCreds := getAppUserCreds(appName, requestUsername)

		if appUserCreds == (config.AuthUserConfig{}) {
			respondWithError(401, "Unauthorized: Invalid Username or Password", c)

			return
		}

		match := secureCompare(appUserCreds.UserName, requestUsername) &&
			secureCompare(appUserCreds.Password, requestPassword)

		if !match {
			respondWithError(401, "Unauthorized: Invalid Username or Password", c)

			return
		}

		c.Next()
	}
}

// getAppUserCreds: It returns the authPair by getting details from config
func getAppUserCreds(appName string, requestUserName string) config.AuthUserConfig {
	authUserValues := appConfig().AuthUser

	switch appName {
	case constants.APIAuthUserFieldName:
		return authUserValues.API
	default:
		return config.AuthUserConfig{}
	}
	return config.AuthUserConfig{}
}

// secureCompare: This function compares two values for their equality
func secureCompare(expected, actual string) bool {
	if subtle.ConstantTimeEq(int32(len(expected)), int32(len(actual))) == 1 {
		return subtle.ConstantTimeCompare([]byte(expected), []byte(actual)) == 1
	}

	return false
}

// isAuthorized: A route is always tied to an app. Every app has a
// username given to it. This function checks whether the authorized app
// only is trying to access this route. It knows which app is trying to
// access this route by parsing the request username sent in basic auth.
func isAuthorized(routeApp []string, requestUsername string) (string, bool) {
	requestAppName := getAppNameFromUsername(requestUsername)
	return utils.Find(routeApp, requestAppName)
}

// getAppNameFromUsername: It returns the auth user based on first string in username
func getAppNameFromUsername(userName string) string {
	return strings.SplitN(userName, "_", 2)[0]
}
