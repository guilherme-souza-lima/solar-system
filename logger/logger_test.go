package logger_test

import (
	"github.com/guilherme-souza-lima/solar-system/entity"
	"github.com/guilherme-souza-lima/solar-system/logger"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	date := time.Now()
	mapping := entity.History{
		URL:        "/user/login",
		StatusCode: 200,
		Message:    "success",
		CreateAt:   date,
		User: entity.User{
			ID:   "123-456-789",
			Name: "Jonny McLane",
			Nick: "JonnyML",
		},
	}
	log := logger.GetInstance().Logger
	log.Warn("[TEST] - URL:" + mapping.URL + ", Message: " + mapping.Message + ", UserID: " + mapping.User.ID)

	//assert.Equal(t, "", log.Exit)
	assert.Equal(t, 200, mapping.StatusCode)
}
