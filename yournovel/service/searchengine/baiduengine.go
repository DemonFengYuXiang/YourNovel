package searchengine

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/url"
	"strings"
	"sync"
	"yournovel/yournovel/conf"
	"yournovel/yournovel/model"
)

type BaiDuSearchEngine struct {
	parseRule       string
	searchRule      string
	domain          string
	parseResultFunc func(searchResult *model.SearchResult)
}

func NewBaiDuSearchEngine(parseResultFunc func(result *model.SearchResult)) *BaiDuSearchEngine {
	return &BaiDuSearchEngine{
		parseRule:       "#content_left h3.t a",
		searchRule:      "intitle:%s 小说 阅读",
		domain:          "http://www.baidu.com/s?wd=%s&ie=utf-8&rn=15&vf_bl=1",
		parseResultFunc: parseResultFunc,
	}
}

func (engine *BaiDuSearchEngine) EngineRun(novelName string, group *sync.WaitGroup) {
	defer group.Done()
	searchKey := url.QueryEscape(fmt.Sprintf(engine.searchRule, novelName))
	requestUrl := fmt.Sprintf(engine.domain, searchKey)

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnHTML(engine.parseRule, func(element *colly.HTMLElement) {
		group.Add(1)
		go engine.extractData(element, group)
	})

	err := c.Visit(requestUrl)
	if err != nil {
		fmt.Println(err)
	}
}

func (engine *BaiDuSearchEngine) extractData(element *colly.HTMLElement, group *sync.WaitGroup) {
	defer group.Done()
	href := element.Attr("href")
	title := element.Text

	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnResponse(func(response *colly.Response) {
		realUrl := response.Request.URL.String()
		isContain := strings.Contains(realUrl, "baidu")
		if isContain {
			return
		}

		host := response.Request.URL.Host
		_,ok := conf.RuleConfig.IgnoreDomain[host]
		if ok {
			return
		}

		isParse := engine.CheckIsParse(host)
		result := &model.SearchResult{Href:realUrl, Title:title, IsParse:isParse, Host:host}
		engine.parseResultFunc(result)
	})

	err := c.Visit(href)
	if err != nil {
		fmt.Println(err)
	}
}

func (engine *BaiDuSearchEngine) CheckIsParse(host string) int {
	isParse := 0
	for key := range conf.RuleConfig.Rule {
		if host == key {
			isParse = 1
			break
		}
	}
	return isParse
}
