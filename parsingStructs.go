package main

import (
	"time"
)

type Webhook struct {
	Commits []struct {
		ID     string `json:"id"`
		Author struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Username string `json:"username"`
		} `json:"author"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
	} `json:"commits"`
	Repository struct {
		ID       int64       `json:"id"`
		Name     string      `json:"name"`
		Language interface{} `json:"language"`
		Owner    struct {
			Name string `json:"name"`
		} `json:"owner"`
	} `json:"repository"`
}
