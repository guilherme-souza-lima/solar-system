package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

var lock = &sync.Mutex{}

type Single struct {
	Logger *logrus.Logger
}

var loggerInstance *Single

func GetInstance() *Single {
	if loggerInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if loggerInstance == nil {
			loggerInstance = &Single{createLogger()}
		}
	}
	return loggerInstance
}

func createLogger() *logrus.Logger {
	textFormatter := &logrus.TextFormatter{
		ForceColors:               false,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              true,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	}
	jsonFormatter := &logrus.JSONFormatter{
		TimestampFormat:   "",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	}

	var log = logrus.New()
	log.Formatter = jsonFormatter
	log.Formatter = textFormatter
	log.Level = logrus.TraceLevel
	log.Out = os.Stdout
	//log.SetReportCaller(true)
	return log
}
