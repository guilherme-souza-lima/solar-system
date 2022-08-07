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
			ID:   "39473b0d-93bf-4309-996a-2b035deb5d6e",
			Name: "Emily Sanchez",
			Nick: "Martramed",
		},
	})

	insert.LoggerElasticsearchPatterTest(entity.History{
		URL:        "/url/logger",
		StatusCode: 400,
		Message:    "error",
		CreateAt:   time.Now(),
		User: entity.User{
			ID:   "18ff9fd7-a675-4a5e-8307-4ddb8b783179",
			Name: "Brandon Bishop",
			Nick: "Subjecome",
		},
	})
}
