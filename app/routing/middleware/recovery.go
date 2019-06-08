package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"secret-server/app/config"
	"secret-server/app/constants"
	"secret-server/app/crerrors"
	"secret-server/app/logger"
	"secret-server/app/routing/response"
	"secret-server/app/utils"
)

// Recovery : Middlewares to catch uncaught rzperror
// This will catch all un-caught panics (if any)
// In case of panic default error response will be set
func Recovery(ctx *gin.Context) {
	defer func(ctx *gin.Context) {
		if rec := recover(); rec != nil {
			err := utils.GetError(rec)
			if config.GetConfig().Application.Mode == gin.DebugMode {
				fmt.Println(err.Error() + string(debug.Stack()))
			} else {
				logger.Get(ctx).WithError(err).Error(crerrors.InternalServerErrorCode)
			}
			ierr := crerrors.NewCrError(ctx, crerrors.InternalServerErrorCode, err)
			responseStruct := response.
				NewResponse(ctx).
				SetError(ierr)
			ctx.Set(constants.Response, responseStruct)
		}
	}(ctx)
	ctx.Next()
}
