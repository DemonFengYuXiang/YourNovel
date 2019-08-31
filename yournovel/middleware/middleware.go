package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MyMiddlewareOption struct {
	IsAuth bool
}

func RequestMiddlewareWrapper(f func(ctx *gin.Context), option MyMiddlewareOption) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		if option.IsAuth {
			err := auth(ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code": 0,
					"msg":  "鉴权失败",
				})
				return
			}
		}
		f(ctx)
	}
}

// 鉴权
func auth(ctx *gin.Context) error {
	return errors.New("鉴权失败")
}
