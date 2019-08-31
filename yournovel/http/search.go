package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
	"yournovel/yournovel/conf"
	"yournovel/yournovel/db/redis"
	"yournovel/yournovel/model"
	"yournovel/yournovel/service/searchengine"
	"yournovel/yournovel/tool"
)

type EngineSearch interface {
	EngineRun(string, *sync.WaitGroup)
}

// 小说搜索
func novelSearch(c *gin.Context) {
	wd, exist := c.GetQuery("wd")
	if !exist || len(wd) == 0 {
		c.Redirect(http.StatusMovedPermanently, conf.Scheme + tool.GetHost() + "/")
		return
	}

	start := time.Now().UnixNano() / 1e6
	results := startEngine(wd)
	end := time.Now().UnixNano() / 1e6

	elapsedTime := (end - start) / 1e3
	searchCount := len(results)
	c.HTML(http.StatusOK, "search_index.html", gin.H{
		"list": results,
		"novelName": wd,
		"elapsedTime": elapsedTime,
		"count": searchCount,
		"head": "search_head",
	})
}

// 开始引擎搜索
func startEngine(novelName string) []*model.SearchResult {

	results, err := redis.RedisConnector.SearchNovelByNovelName(novelName)
	if err == nil {
		return results
	}

	group := sync.WaitGroup{}
	results = make([]*model.SearchResult, 0)
	for _,engine := range conf.RuleConfig.Engines {
		var searchEngine EngineSearch
		group.Add(1)
		switch engine {
		case "baidu":
			searchEngine = searchengine.NewBaiDuSearchEngine(func(result *model.SearchResult) {
				results = append(results, result)
			})
		}

		if searchEngine != nil {
			go searchEngine.EngineRun(novelName, &group)
		}
	}
	group.Wait()
	if len(results) > 0 {
			err = redis.RedisConnector.SaveSearchResultByNovelName(novelName, results)
		if err != nil {
			fmt.Printf("Redis Save Failure %s", err)
		}
	}

	return results
}
