package messaging

import (
	b "bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	result "github.com/heaptracetechnology/microservice-telegram/result"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

type BotMessage struct {
	ChatID      int64  `json:"chat_id"`
	Message     string `json:"message"`
	Username    string `json:"username"`
	ImageBase64 string `json:"image"`
	URL         string `json:"url"`
}

type Subscribe struct {
	Data      Data   `json:"data"`
	Offset    int    `json:"offset"`
	Endpoint  string `json:"endpoint"`
	Id        string `json:"id"`
	IsTesting bool   `json:"istesting"`
}

type Data struct {
	Channel string `json:"channel"`
}

type UpdateResponse struct {
	UpdateId string `json:"update_id"`
	Channel  string `json:"channel"`
	Id       string `json:"id"`
	Pattern  string `json:"pattern"`
	Direct   string `json:"direct"`
}

type Payload struct {
	EventId     string          `json:"eventID"`
	EventType   string          `json:"eventType"`
	ContentType string          `json:"contentType"`
	Data        tgbotapi.Update `json:"data"`
}

var Listner = make(map[string]Subscribe)
var rtmstarted bool
var offset string
var isBotRunning bool
var bot *tgbotapi.BotAPI

//Get Bot Details
func GetBotDetails(responseWriter http.ResponseWriter, request *http.Request) {

	var botToken = os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	botDetails, _ := bot.GetMe()

	bytes, _ := json.Marshal(botDetails)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Send Message By Bot
func SendMessage(responseWriter http.ResponseWriter, request *http.Request) {

	var botToken = os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
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

	var botToken = os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
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

//Get Chat
func GetChat(responseWriter http.ResponseWriter, request *http.Request) {

	var botToken = os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
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

	var botToken = os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
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

	var botToken = os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	body, bodyErr := ioutil.ReadAll(request.Body)
	if bodyErr != nil {
		result.WriteErrorResponse(responseWriter, bodyErr)
		return
	}

	defer request.Body.Close()

	var botMessage BotMessage
	unmarshalErr := json.Unmarshal(body, &botMessage)
	if unmarshalErr != nil {
		result.WriteErrorResponse(responseWriter, unmarshalErr)
		return
	}

	data, decodeStringErr := base64.StdEncoding.DecodeString(botMessage.ImageBase64)
	if decodeStringErr != nil {
		result.WriteErrorResponse(responseWriter, decodeStringErr)
		return
	}

	filepath := os.TempDir() + "/" + time.Now().String() + ".jpeg"
	f, createErr := os.Create(filepath)
	if createErr != nil {
		result.WriteErrorResponse(responseWriter, createErr)
		return
	}

	defer f.Close()

	if _, err := f.Write(data); err != nil {
		panic(err)
	}
	if err := f.Sync(); err != nil {
		panic(err)
	}

	uploadphoto := tgbotapi.NewPhotoUpload(botMessage.ChatID, filepath)
	reponse, errr := bot.Send(uploadphoto)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}
	bytes, _ := json.Marshal(reponse)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Subscribe
func SubscribeUpdate(responseWriter http.ResponseWriter, request *http.Request) {

	//isBotRunning := false

	if isBotRunning != true {
		var botToken = os.Getenv("BOT_TOKEN")
		bot, _ = tgbotapi.NewBotAPI(botToken)
		isBotRunning = true
	}

	decoder := json.NewDecoder(request.Body)

	var listner Subscribe
	errr := decoder.Decode(&listner)
	if errr != nil {
		result.WriteErrorResponse(responseWriter, errr)
		return
	}

	res2, _ := json.Marshal(listner)
	fmt.Println(string(res2))

	Listner[listner.Id] = listner
	if rtmstarted != true {
		go TeleGramRTM()
		rtmstarted = true
	}

	bytes, _ := json.Marshal("Subscribed")
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//Unsubscribe
func UnsubscribeUpdate(responseWriter http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)
	var id string
	err := decoder.Decode(&id)
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}
	if len(Listner) > 0 {
		delete(Listner, id)
	}

	bytes, _ := json.Marshal("UnSubscribed")
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

func TeleGramRTM() {
	istest := false
	quit := make(chan struct{})
	for {
		if len(Listner) > 0 {
			for k, v := range Listner {
				go getMessageUpdates(k, v)
				istest = v.IsTesting
			}
		} else {
			rtmstarted = false
			break
		}
		time.Sleep(3 * time.Second)
		if istest {
			close(quit)
			break
		}
	}
}

func getMessageUpdates(userid string, sub Subscribe) {

	hc := http.Client{}
	var param tgbotapi.UpdateConfig
	if sub.Offset > 0 {
		param.Offset = sub.Offset
	}

	getUpdates, updateErr := bot.GetUpdates(param)
	if updateErr != nil {
		fmt.Println("updateErr :", updateErr)
	}

	var messages []tgbotapi.Update

	messages = getUpdates
	var newMsg tgbotapi.Update

	for _, msg := range messages {
		newMsg = msg

	}

	var response Payload

	response.ContentType = "application" + "/" + "json"
	response.EventType = "hears"
	response.EventId = sub.Id
	response.Data = newMsg

	requestBody := new(b.Buffer)
	err := json.NewEncoder(requestBody).Encode(response)
	if err != nil {
		fmt.Println(" request err :", err)
	}

	if newMsg.UpdateID != sub.Offset && newMsg.ChannelPost.Chat.UserName == sub.Data.Channel {

		req, errr := http.NewRequest("POST", sub.Endpoint, requestBody)
		if errr != nil {
			fmt.Println(" request err :", errr)
		}
		_, err := hc.Do(req)
		if err != nil {
			fmt.Println("Client error", err)
		}
		if newMsg.UpdateID > sub.Offset {
			sub.Offset = newMsg.UpdateID
		}

		Listner[sub.Id] = sub
	}
}
