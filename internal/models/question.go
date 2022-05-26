package models

type Question struct {
	Id            int64             `json:"id"`
	Description   string            `json:"description"`
	Answers       map[string]string `json:"answers"`
	CorrectAnswer string            `json:"correct_answer"`
}
