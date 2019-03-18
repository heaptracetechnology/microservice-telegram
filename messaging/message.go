package messaging

import (
	"encoding/json"
	result "github.com/heaptracetechnology/microservice-telegram/result"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"net/http"
	"os"
)

type BotMessage struct {
	ChatID   int64  `json:"chat_id"`
	Message  string `json:"message"`
	Username string `json:"username"`
}

//Get Bot Details
func GetBotDetails(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	bot, err := tgbotapi.NewBotAPI(accessToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
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

//Send Group Message By Bot
func SendGroupMessage(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	bot, err := tgbotapi.NewBotAPI(accessToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
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
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	bytes, _ := json.Marshal(res)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Send Channel Message By Bot
func SendChannelMessage(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	bot, err := tgbotapi.NewBotAPI(accessToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var param *BotMessage
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	newChannelMessage := tgbotapi.NewMessageToChannel(param.Username, param.Message)
	response, err := bot.Send(newChannelMessage)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	bytes, _ := json.Marshal(response)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

// //Get Chat
func GetChat(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	bot, err := tgbotapi.NewBotAPI(accessToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	var param BotMessage
	decoder := json.NewDecoder(request.Body)
	DecodeErr := decoder.Decode(&param)
	if DecodeErr != nil {
		result.WriteErrorResponse(responseWriter, DecodeErr)
		return
	}

	var chatConfig tgbotapi.ChatConfig
	chatConfig.ChatID = param.ChatID

	getChatDetails, err := bot.GetChat(chatConfig)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	bytes, _ := json.Marshal(getChatDetails)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Leave Chat
func LeaveChat(responseWriter http.ResponseWriter, request *http.Request) {

	var accessToken = os.Getenv("ACCESS_TOKEN")

	bot, err := tgbotapi.NewBotAPI(accessToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	var param BotMessage
	decoder := json.NewDecoder(request.Body)
	DecodeErr := decoder.Decode(&param)
	if DecodeErr != nil {
		result.WriteErrorResponse(responseWriter, DecodeErr)
		return
	}

	var chatConfig tgbotapi.ChatConfig
	chatConfig.ChatID = param.ChatID

	leaveChat, err := bot.LeaveChat(chatConfig)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	bytes, _ := json.Marshal(leaveChat)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
