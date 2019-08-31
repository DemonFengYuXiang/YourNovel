package redis

import (
	"fmt"
	"testing"
	"yournovel/yournovel/model"
)

func TestRedisClinet(t *testing.T) {

	InitRedisClient()

	searchResult := &model.SearchResult{}
	searchResult.Title = "帝霸"
	searchResult.Host = "hao123.com"
	searchResult.IsParse = 1
	searchResult.Href = "http://www.baidu.com"

	searchResult2 := &model.SearchResult{}
	searchResult2.Title = "帝霸"
	searchResult2.Host = "hao123.com"
	searchResult2.IsParse = 1
	searchResult2.Href = "http://www.baidu.com"

	searchResults := []*model.SearchResult{
		searchResult,
		searchResult2,
	}

	err := RedisConnector.SaveSearchResultByNovelName("帝霸", searchResults)
	if err != nil {
		t.Errorf("%s", err)
	}

	results, err := RedisConnector.SearchNovelByNovelName("帝霸")
	if err != nil {
		t.Errorf("%s", err)
	}

	for _,result := range results {
		fmt.Println(result.Href)
	}
}
