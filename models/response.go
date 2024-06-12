package models

type Response struct {
	Message  string   `json:"message"`
	Employee Employee `json:"employee"`
}
