# solar-system

### Logger

`log := logger.GetInstance().Logger`

`log.Info("success")`

`log.Error("error")`

### Elastic

`brokers := []string{"http: //localhost:9200", "http: //localhost:9201"}`

`insert := elastic.NewLoggerElastic("logger_ss", brokers)`

`insert.LoggerElasticsearch(entity.MappingElastic{...})`
