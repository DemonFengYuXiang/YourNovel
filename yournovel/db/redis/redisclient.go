package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"time"
	"yournovel/yournovel/conf"
	"yournovel/yournovel/model"
)

type RedisClient struct {
	client *redis.Client
}

var RedisConnector *RedisClient

func InitRedisClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     conf.RedisAddr,
		Password: conf.RedisPassword,
		DB:       conf.RedisDb,
	})

	RedisConnector = &RedisClient{
		client: client,
	}

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)
	return
}

// 根据小说名称将搜索结果保存至Redis中
func (connector *RedisClient) SaveSearchResultByNovelName(novelName string, searchResult []*model.SearchResult) error {
	if searchResult == nil {
		return errors.New(fmt.Sprintf("novel:%s searchResult failure", novelName))
	}

	searchJson, err := jsoniter.ConfigFastest.MarshalToString(searchResult)
	if err != nil {
		return err
	}

	err = connector.client.Set(novelName, searchJson, 24 * time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

// 从Redis中获取搜索小说
func (connector *RedisClient) SearchNovelByNovelName(novelName string) ([]*model.SearchResult, error) {
	var searchResult []*model.SearchResult
	searchJson, err := connector.client.Get(novelName).Result()
	if err == redis.Nil {
		return searchResult, errors.New(fmt.Sprintf("Redis Error: %s not exists??", novelName))
	}

	err = jsoniter.ConfigFastest.UnmarshalFromString(searchJson, &searchResult)
	if err != nil {
		return nil, err
	}

	return searchResult, nil
}
