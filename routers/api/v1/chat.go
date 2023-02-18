package v1

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/service/chat_service"
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