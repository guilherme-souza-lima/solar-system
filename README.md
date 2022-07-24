# solar-system

### Log

`log := logger.GetInstance().Logger`

`og.Info("success")`

`log.Error("error")`

### Elastic

`brokers := []string{"http: //localhost:9200", "http: //localhost:9201"}`

`insert := elastic.NewLoggerElastic("logger_ss", brokers)`

`insert.LoggerElasticsearch(entity.MappingElastic{...})`