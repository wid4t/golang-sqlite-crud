package models

type Employee struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	JobName string `json:"jobName,omitempty"`
}
