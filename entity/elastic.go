package entity

import "time"

type MappingElastic struct {
	StatusCode int
	Level      string
	Message    MessageElastic
	Date       time.Time
	User       UserElastic
}

type MessageElastic struct {
	Message string
	Local   string
}

type UserElastic struct {
	ID string
}
