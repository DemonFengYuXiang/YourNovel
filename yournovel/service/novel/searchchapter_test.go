package novel

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"testing"
)

func TestSearchChapter(t *testing.T) {

	c := colly.NewCollector(
		colly.DetectCharset(),
		)
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnHTML("#at", func(element *colly.HTMLElement) {
	})

	c.Visit("https://www.x23us.com/html/64/64889/")
}
