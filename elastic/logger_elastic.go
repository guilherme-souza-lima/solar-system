package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
	"github.com/guilherme-souza-lima/solar-system/entity"
	"github.com/guilherme-souza-lima/solar-system/logger"
)

type LoggerElastic struct {
	Index   string
	Brokers []string
}

func NewLoggerElastic(Index string, Broker []string) LoggerElastic {
	return LoggerElastic{Index, Broker}
}

func (l LoggerElastic) LoggerElasticsearch(mapping entity.History) error {
	log := logger.GetInstance().Logger
	log.Warn("Init LoggerElasticsearch")

	client, err := GetESClient(l.Brokers)
	if err != nil {
		log.Error("Error MyLib: " + err.Error())
		return err
	}

	dataJSON, err := json.Marshal(mapping)
	if err != nil {
		log.Error("Error MyLib: " + err.Error())
		return err
	}

	id := uuid.New()
	req := esapi.IndexRequest{
		Index:      l.Index,
		DocumentID: id.String(),
		Body:       bytes.NewReader(dataJSON),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		log.Error("Error MyLib: " + err.Error())
		log.Error("Logger Elasticsearch: " + l.Index)
		return err
	}
	defer res.Body.Close()

	log.Info("Logger Elasticsearch: " + l.Index)
	log.Warn("End LoggerElasticsearch")
	return nil

}
