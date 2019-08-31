package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 首页
func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"head": "index_head",
	})
}
