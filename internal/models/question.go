package models

type Question struct {
	ID            int64             `json:"id"`
	Description   string            `json:"description"`
	Answers       map[string]string `json:"answers"`
	CorrectAnswer string            `json:"-"`
}
