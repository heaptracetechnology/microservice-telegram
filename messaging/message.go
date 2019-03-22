package messaging

import (
	//"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	result "github.com/heaptracetechnology/microservice-telegram/result"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	//"path/filepath"
	//"bytes"
	//"io"
)

type BotMessage struct {
	ChatID      int64  `json:"chat_id"`
	Message     string `json:"message"`
	Username    string `json:"username"`
	ImageBase64 string `json:"image"`
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
func SendMessage(responseWriter http.ResponseWriter, request *http.Request) {

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

//Send Photo
func SendPhoto(responseWriter http.ResponseWriter, request *http.Request) {

	fmt.Println("===============")
	var accessToken = os.Getenv("ACCESS_TOKEN")

	bot, err := tgbotapi.NewBotAPI(accessToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println("======errr=======")
	}

	defer request.Body.Close()

	var botMessage BotMessage
	er := json.Unmarshal(body, &botMessage)
	if er != nil {
		fmt.Println("======errr1=======")
	}
	fmt.Println("image.ImageBase64 ::: ", botMessage.ImageBase64)
	data, err := base64.StdEncoding.DecodeString(botMessage.ImageBase64)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	filepath := os.TempDir() + "/" + time.Now().String() + ".jpeg"
	f, err1 := os.Create(filepath)
	if err1 != nil {
		fmt.Println("errr file create  :::: ", err1)
		panic(err)

	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	leaveChat := tgbotapi.NewPhotoUpload(botMessage.ChatID, filepath)
	reponse, errr := bot.Send(leaveChat)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}
	bytes, _ := json.Marshal(reponse)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
