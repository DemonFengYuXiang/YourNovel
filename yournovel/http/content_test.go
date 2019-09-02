package http

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
	"testing"
	"yournovel/yournovel/tool"
)

func TestNextChapter(t *testing.T) {

	c := colly.NewCollector(
		colly.DetectCharset(),
		)

	html := ""
	c.OnResponse(func(response *colly.Response) {
		html = string(response.Body)
	})

	c.Visit("https://www.qb5200.tw/xiaoshuo/40/40994/13951060.html")

	html = strings.ReplaceAll(html, "<<", "")
	html = strings.ReplaceAll(html, ">>", "")

	// 参考https://greasyfork.org/zh-CN/scripts/292-my-novel-reader
	next_reg := `(<a\s+.*?>.*[第上前下后][一]?[0-9]{0,6}?[页张个篇章节步].*?</a>)`
    judge_reg := `[第上前下后][一]?[0-9]{0,6}?[页张个篇章节步]`

    re := regexp.MustCompile(next_reg)
    allFind := re.FindAll([]byte(html), -1)
    allByte := bytes.Join(allFind, []byte("\n"))
    reader := bytes.NewReader(allByte)
    doc,_ := goquery.NewDocumentFromReader(reader)

    judgeRe := regexp.MustCompile(judge_reg)

    doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		text = strings.ReplaceAll(text, " ", "")
		isNext := judgeRe.MatchString(text)
		if isNext {
			href,exist := selection.Attr("href")
			if !exist {
				return
			}

			url := tool.UrlJoin(href, "http://www.hao123.com")
			fmt.Println(url)
		}
	})
}

func TestTtile(t *testing.T) {

	c := colly.NewCollector(
		colly.DetectCharset(),
	)

	html := ""
	c.OnResponse(func(response *colly.Response) {
		html = string(response.Body)
	})

	c.Visit("https://www.qb5200.tw/xiaoshuo/40/40994/13951060.html")

	titleReg := `(第?\s*[一二两三四五六七八九十○零百千万亿0-9１２３４５６７８９０]{1,6}\s*[章回卷节折篇幕集]\s*.*?)[_,-]`
	re := regexp.MustCompile(titleReg)
	titleRes := re.FindAllString(html, -1)
	title := ""
	reader := strings.NewReader(html)
	doc, _ := goquery.NewDocumentFromReader(reader)
	if titleRes != nil {
		title = titleRes[1]
	}else {
		doc.Find("h1").Each(func(i int, selection *goquery.Selection) {
			if i == 0 {
				title = selection.Text()
			}
			return
		})
	}

	if len(title) <= 0 {
		doc.Find("title").Each(func(i int, selection *goquery.Selection) {
			if i == 0 {
				title = selection.Text()
			}
			return
		})
	}

	fmt.Println(title)
}

