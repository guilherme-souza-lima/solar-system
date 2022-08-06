package solar_system

import (
	"github.com/guilherme-souza-lima/solar-system/elastic"
	"github.com/guilherme-souza-lima/solar-system/entity"
	"time"
)

func main() {
	brokers := []string{"", ""}
	insert := elastic.NewLoggerElastic("", brokers)
	insert.LoggerElasticsearch(entity.MappingElastic{
		StatusCode: 0,
		Level:      "",
		Message:    entity.MessageElastic{},
		Date:       time.Time{},
		User:       entity.UserElastic{},
	})
}
