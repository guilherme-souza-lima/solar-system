package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/google/uuid"
	"solar-system/entity"
	"solar-system/logger"
)

type LoggerElastic struct {
	Index   string
	Brokers []string
}

func NewLoggerElastic(Index string, Broker []string) LoggerElastic {
	return LoggerElastic{Index, Broker}
}

func (l LoggerElastic) LoggerElasticsearch(mapping entity.MappingElastic) error {
	log := logger.GetInstance().Logger

	client, err := GetESClient(l.Brokers)
	if err != nil {
		log.Error(err)
		return err
	}

	dataJSON, err := json.Marshal(mapping)
	if err != nil {
		log.Error(err)
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
		log.Error(err)
		return err
	}
	defer res.Body.Close()

	log.Info("success")
	return nil
}
