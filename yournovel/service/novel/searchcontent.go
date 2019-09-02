package novel

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"html/template"
	"net/url"
	"regexp"
	"strings"
	"yournovel/yournovel/conf"
	"yournovel/yournovel/fetcher"
	"yournovel/yournovel/model"
	"yournovel/yournovel/tool"
)

func SearchContentOfNovel(contentUrl string) (*model.NovelContent, error) {
	c := fetcher.NewNovelConnector()
	var novelContent model.NovelContent

	requestURI, err := url.ParseRequestURI(contentUrl)
	if err != nil {
		return &novelContent, err
	}

	host := requestURI.Host
	contentSelector,ok := conf.RuleConfig.Rule[host]["content_selector"].(string)
	if !ok {
		return &novelContent, err
	}

	c.OnResponse(func(response *colly.Response) {
		webHtml := string(response.Body)
		chapterUrls := getPreChapterAndNextChapterFromHtml(webHtml, contentUrl)
		novelContent.PreChapter = chapterUrls[0]
		novelContent.NextChapter = chapterUrls[1]

		title := getTitleFromHtml(webHtml)
		novelContent.Title = title
	})

	c.OnHTML(contentSelector, func(element *colly.HTMLElement) {
		html,err := element.DOM.Html()
		if err != nil {
			return
		}
		novelContent.Content = template.HTML(html)
	})

	err = c.Visit(contentUrl)
	if err != nil {
		return &novelContent, err
	}

	return &novelContent, nil
}

// 从HTML文件中获得上一页以及下一页
func getPreChapterAndNextChapterFromHtml(html string, url string) []string {
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

	chapterUrls := make([]string, 0, 2)
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		text = strings.ReplaceAll(text, " ", "")
		isNext := judgeRe.MatchString(text)
		if isNext {
			href,exist := selection.Attr("href")
			if !exist {
				return
			}

			chapterUrl := tool.UrlJoin(href, url)
			chapterUrls = append(chapterUrls, chapterUrl)
		}
	})
	return chapterUrls
}

func getTitleFromHtml(html string) string {
	titleReg := `(第?\s*[一二两三四五六七八九十○零百千万亿0-9１２３４５６７８９０]{1,6}\s*[章回卷节折篇幕集]\s*.*?)[_,-]`
	re := regexp.MustCompile(titleReg)
	titleRes := re.FindAllString(html, 0)
	title := ""
	reader := strings.NewReader(html)
	doc, _ := goquery.NewDocumentFromReader(reader)
	if titleRes != nil {
		title = titleRes[0]
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

	return title
}

