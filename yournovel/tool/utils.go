package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yournovel/yournovel/conf"
)

func GetHost() string {
	host := conf.Host+":"+ strconv.Itoa(conf.Port)
	return host
}

func SuccessResponse(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg": msg,
		"data": data,
	})
}

func ErrorResponse(ctx *gin.Context, msg string, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": msg,
		"data": data,
	})
}


