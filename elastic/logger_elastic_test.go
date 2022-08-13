package elastic_test

import (
	"github.com/google/uuid"
	"github.com/guilherme-souza-lima/solar-system/elastic"
	"github.com/guilherme-souza-lima/solar-system/entity"
	"testing"
	"time"
)

func user(url, message, name, nick string, status int) entity.History {
	id := uuid.New()
	return entity.History{
		URL:        url,
		StatusCode: status,
		Message:    message,
		CreateAt:   time.Now(),
		User: entity.User{
			ID:   id.String(),
			Name: name,
			Nick: nick,
		},
	}
}

func TestLoggerElasticSearch(t *testing.T) {
	//brokers := []string{"http://167.114.147.66:9200", "http://167.114.147.66:9201"}
	brokers := []string{"http://localhost:9200", "http://localhost:9201"}
	url := "/user/login"

	insert := elastic.NewLoggerElastic("logger_ss", brokers)

	insert.LoggerElasticsearch(user(url, "", "", "", 200))

	insert.LoggerElasticsearch(entity.History{
		URL:        "/user/login",
		StatusCode: 200,
		Message:    "success",
		CreateAt:   time.Now(),
		User: entity.User{
			ID:   "39473b0d-93bf-4309-996a-2b035deb5d6e",
			Name: "Emily Sanchez",
			Nick: "Martramed",
		},
	})

	insert.LoggerElasticsearch(entity.History{
		URL:        "/user/login",
		StatusCode: 400,
		Message:    "error",
		CreateAt:   time.Now(),
		User: entity.User{
			ID:   "18ff9fd7-a675-4a5e-8307-4ddb8b783179",
			Name: "Brandon Bishop",
			Nick: "Subjecome",
		},
	})

	insert.LoggerElasticsearch(entity.History{
		URL:        "/user/login",
		StatusCode: 200,
		Message:    "success",
		CreateAt:   time.Now(),
		User: entity.User{
			ID:   "",
			Name: "",
			Nick: "",
		},
	})
}
