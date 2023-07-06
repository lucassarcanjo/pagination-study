package model

import "time"

type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name" fake:"{firstname} {lastname}"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Company     string    `json:"company"`
	JobTitle    string    `json:"job_title"`
	CreatedTime time.Time `json:"created_time"`
}
