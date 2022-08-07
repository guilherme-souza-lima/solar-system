package entity

import "time"

type History struct {
	URL        string    `json:"url"`
	StatusCode int       `json:"status_code"`
	Message    string    `json:"message"`
	CreateAt   time.Time `json:"create_at"`
	User       User      `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Nick string `json:"nick"`
}
