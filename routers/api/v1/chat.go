package v1

import (
	"github.com/AlmazDefourten/digital-helper/pkg/app"
	"github.com/AlmazDefourten/digital-helper/pkg/e"
	"github.com/AlmazDefourten/digital-helper/service/chat_service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAnswer(c *gin.Context) {
	appG := app.Gin{C: c}
	question := c.Query("question")

	chatService := chat_service.AnswerDto{}

	answer, err := chatService.GetAnswer(question)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, answer)
}

func StartChat(c *gin.Context) {
	appG := app.Gin{C: c}

	chatService := chat_service.Chat{}

	answer, err := chatService.StartChat()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, answer)
}
