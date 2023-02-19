package chat_service

import "github.com/AlmazDefourten/digital-helper/models"

type Chat struct {
}

func (chat *Chat) StartChat() (*models.Answer, error) {
	var answer = models.Answer{}
	answ, err := models.Start()
	answer.Answers = answ
	if err != nil {
		return nil, err
	}

	return &answer, nil
}
