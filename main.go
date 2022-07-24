package main

import (
	"github.com/google/uuid"
	"solar-system/elastic"
	"solar-system/entity"
	"time"
)

func main() {
	id := uuid.New()
	mapping := entity.MappingElastic{
		StatusCode: 200,
		Level:      "INFO",
		Message: entity.MessageElastic{
			Message: "Message, info or error",
			Local:   "API",
		},
		Date: time.Now(),
		User: entity.UserElastic{
			ID: id.String(),
		},
	}

	brokers := []string{"http: //localhost:9200", "http: //localhost:9201"}
	insert := elastic.NewLoggerElastic("logger_ss", brokers)
	insert.LoggerElasticsearch(mapping)

}
