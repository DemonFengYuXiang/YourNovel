package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yournovel/yournovel/service/novel"
)

func novelContent(c *gin.Context) {

	chapterUrl, exist := c.GetQuery("chapter_url")
	if !exist {
		c.Redirect(http.StatusMovedPermanently, "/")
		return
	}

	contentUrl, exist := c.GetQuery("content_url")
	if !exist {
		c.Redirect(http.StatusMovedPermanently, chapterUrl)
		return
	}
	novelName, exist := c.GetQuery("novel_name")
	if !exist {
		c.Redirect(http.StatusMovedPermanently, chapterUrl)
		return
	}

	contentTitle, exist := c.GetQuery("content_title")
	if !exist {
		c.Redirect(http.StatusMovedPermanently, chapterUrl)
		return
	}

	content, err := novel.SearchContentOfNovel(contentUrl)
	if err != nil {
		fmt.Println(content)
		//c.Redirect(http.StatusMovedPermanently, chapterUrl)
		return
	}
	content.NovelName = novelName
	content.ContentUrl = contentUrl

	c.HTML(http.StatusOK, "content_index.html", gin.H{
		"content": content,
		"chapter_url": chapterUrl,
		"novel_name": novelName,
		"content_title": contentTitle,
		"head": "content_head",
	})
}
