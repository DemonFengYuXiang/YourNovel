package novel

import (
	"github.com/gocolly/colly"
	"html/template"
	"net/url"
	"yournovel/yournovel/conf"
	"yournovel/yournovel/fetcher"
	"yournovel/yournovel/model"
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
