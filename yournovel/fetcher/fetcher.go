package fetcher

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func NewNovelConnector() *colly.Collector {
	c := colly.NewCollector(
		colly.DetectCharset(),
	)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	return c
}
