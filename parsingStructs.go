package main

import (
	"time"
)


<<<<<<< HEAD
type Webhook struct {
        Commits []struct {
                ID     string   `json:"id"`
                Author struct {
                        Email    string `json:"email"`
                        Name     string `json:"name"`
                        Username string `json:"username"`
                } `json:"author"`
                Message   string        `json:"message"`
                Timestamp time.Time     `json:"timestamp"`
        } `json:"commits"`
        Repository struct {
                ID               int64       `json:"id"`
                Name             string      `json:"name"`
                Owner            struct {
                        Name              string `json:"name"`
                } `json:"owner"`
        } `json:"repository"`
}

