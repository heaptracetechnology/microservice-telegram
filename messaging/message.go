package messaging

import (
	"encoding/json"
	"fmt"
	result "github.com/heaptracetechnology/microservice-telegram/result"
	"gopkg.in/telegram-bot-api.v4"
	"net/http"
	"os"
)

type BotMessage struct {
	ChatID  int64  `json:"chat_id"`
	Message string `json:"message"`
}

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

//Send Bot Message
func Send(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	bot, err := tgbotapi.NewBotAPI(accessToken)
	if err != nil {
		fmt.Println("err new bot api : ", err)
	}

	decoder := json.NewDecoder(request.Body)
	var param *BotMessage
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	sendMessage := tgbotapi.NewMessage(param.ChatID, param.Message)
	res, err := bot.Send(sendMessage)
	if err != nil {
		fmt.Println("err new message :", err)
	}

	bytes, _ := json.Marshal(res)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
