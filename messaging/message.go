package messaging

import (
	"encoding/json"
	result "github.com/heaptracetechnology/microservice-telegram/result"
	"gopkg.in/telegram-bot-api.v4"
	"net/http"
	"os"
)

//Get Bot Details
func GetBotDetails(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	bot, erro := tgbotapi.NewBotAPI(accessToken)
	if erro != nil {
		result.WriteErrorResponse(responseWriter, erro)
		return
	}

	botDetails, err := bot.GetMe()
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	bytes, _ := json.Marshal(botDetails)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
