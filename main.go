package main

import (
	"github.com/guilherme-souza-lima/solar-system/elastic"
	"github.com/guilherme-souza-lima/solar-system/entity"
	"time"
)

func main() {
	brokers := []string{"http://167.114.147.66:9200", "http://167.114.147.66:9201"}
	insert := elastic.NewLoggerElastic("logger_ss", brokers)
	insert.LoggerElasticsearchPatterTest(entity.History{
		URL:        "/url/logger",
		StatusCode: 200,
		Message:    "success",
		CreateAt:   time.Now(),
		User: entity.User{
			ID:   "22",
			Name: "33",
			Nick: "44",
		},
	})
}
