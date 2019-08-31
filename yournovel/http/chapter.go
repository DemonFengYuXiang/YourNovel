package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"yournovel/yournovel/service/novel"
	"yournovel/yournovel/tool"
)

func novelChapter(c *gin.Context) {

	webSiteUrl, exist := c.GetQuery("url")
	if !exist {
		tool.ErrorResponse(c, "源网址不存在", webSiteUrl)
		return
	}
	novelName, exist := c.GetQuery("novel_name")
	if !exist {
		c.Redirect(http.StatusMovedPermanently, webSiteUrl)
		return
	}

	novelChapter, err := novel.SearchChapterOfNovel(webSiteUrl, novelName)
	if err != nil {
		//c.Redirect(http.StatusMovedPermanently, webSiteUrl)
		return
	}

	c.HTML(http.StatusOK, "chapter_index.html", gin.H{
		"chapter": novelChapter,
		"head": "chapter_head",
	} )
}
