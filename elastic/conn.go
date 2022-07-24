package elastic

import "github.com/elastic/go-elasticsearch/v7"

func GetESClient(brokers []string) (*elasticsearch.Client, error) {

	cfg := elasticsearch.Config{
		Addresses: brokers,
	}
	return elasticsearch.NewClient(cfg)

}
