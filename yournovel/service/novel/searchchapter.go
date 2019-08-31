package novel

import (
	"fmt"
	"github.com/gocolly/colly"
	"html/template"
	"net/url"
	"yournovel/yournovel/conf"
	"yournovel/yournovel/fetcher"
	"yournovel/yournovel/model"
)

func SearchChapterOfNovel(novelUrl string, novelName string) (*model.NovelChapter, error) {
	c := fetcher.NewNovelConnector()

	var novelChapter model.NovelChapter
	requestURI, err := url.ParseRequestURI(novelUrl)
	if err != nil {
		return &novelChapter, err
	}

	host := requestURI.Host
	chapterSelector,ok := conf.RuleConfig.Rule[host]["chapter_selector"].(string)
	if !ok {
		return &novelChapter, err
	}

	c.OnHTML(chapterSelector, func(element *colly.HTMLElement) {
		html,_ := element.DOM.Parent().Html()
		if err != nil {
			return
		}

		novelChapter.Content = template.HTML(html)
		novelChapter.Name = novelName
		novelChapter.OriginUrl = novelUrl
		novelChapter.LinkPrefix = conf.RuleConfig.Rule[host]["link_prefix"].(string)
		novelChapter.Domain = fmt.Sprintf("%s://%s", requestURI.Scheme, requestURI.Host)
	})

	err = c.Visit(novelUrl)
	if err != nil {
		return &novelChapter, err
	}

	return &novelChapter, nil
}
