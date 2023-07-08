package model

type User struct {
	Base

	Name     string `json:"name" fake:"{name}"`
	Email    string `json:"email" fake:"{email}"`
	Phone    string `json:"phone" fake:"{phone}"`
	Company  string `json:"company" fake:"{company}"`
	JobTitle string `json:"job_title" fake:"{jobtitle}"`
}
