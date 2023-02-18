package chat_service

import "github.com/EDDYCJY/go-gin-example/models"

type AnswerDto struct {
	Question string
	Answer   string
}

func (ans *AnswerDto) GetAnswer(question string) (*models.Answer, error) {
	answer, err := models.GetAnswer(question)
	if err != nil {
		return nil, err
	}

	return answer, nil
}
