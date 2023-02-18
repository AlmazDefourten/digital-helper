package models

import (
	"github.com/jinzhu/gorm"
	"strings"
)

type Answer struct {
	ID        int    `json:"id"`
	Questions string `json:"question"`
	Answers   string `json:"answer"`
}

// GetAnswer get answer for question
func GetAnswer(question string) (*Answer, error) {
	var answer Answer
	err := db.Where("to_tsvector('russian', questions) @@ to_tsquery('russian', ?)", strings.Replace(question, " ", " & ", -1)).First(&answer).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &answer, nil
}
