package messaging

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/cloudevents/sdk-go"
	result "github.com/heaptracetechnology/microservice-telegram/result"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
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

var Listner = make(map[string]Subscribe)
var rtmstarted bool
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

	if !isBotRunning {
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

	Listner[listner.Id] = listner
	if !rtmstarted {
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

	var param tgbotapi.UpdateConfig
	if sub.Offset > 0 {
		param.Offset = sub.Offset
	}

	getUpdates, updateErr := bot.GetUpdates(param)
	if updateErr != nil {
		fmt.Println("updateErr :", updateErr)
	}

	var messages = getUpdates
	var newMsg tgbotapi.Update

	for _, msg := range messages {
		newMsg = msg
	}

	contentType := "application/json"
	s1 := strings.Split(sub.Endpoint, "//")
	_, ip := s1[0], s1[1]
	s := strings.Split(ip, ":")
	_, port := s[0], s[1]
	sub.Endpoint = "http://192.168.0.61:" + string(port)
	t, err := cloudevents.NewHTTPTransport(
		cloudevents.WithTarget(sub.Endpoint),
		cloudevents.WithStructuredEncoding(),
	)
	if err != nil {
		log.Printf("failed to create transport, %v", err)
		return
	}

	c, err := cloudevents.NewClient(t,
		cloudevents.WithTimeNow(),
	)
	if err != nil {
		log.Printf("failed to create client, %v", err)
		return
	}

	source, err := url.Parse(sub.Endpoint)
	event := cloudevents.Event{
		Context: cloudevents.EventContextV01{
			EventID:     sub.Id,
			EventType:   "hears",
			Source:      cloudevents.URLRef{URL: *source},
			ContentType: &contentType,
		}.AsV01(),
		Data: newMsg,
	}

	if newMsg.UpdateID != sub.Offset && newMsg.ChannelPost.Chat.UserName == sub.Data.Channel {

		resp, err := c.Send(context.Background(), event)
		if err != nil {
			log.Printf("failed to send: %v", err)
		}
		if resp != nil {
			fmt.Printf("Response:\n%s\n", resp)
			fmt.Printf("Got Event Response Context: %+v\n", resp.Context)
			data := event
			if err := resp.DataAs(event); err != nil {
				fmt.Printf("Got Data Error: %s\n", err.Error())
			}
			fmt.Printf("Got Response Data: %+v\n", data)
		} else {
			log.Printf("event sent at %s", time.Now())
		}
		if newMsg.UpdateID > sub.Offset {
			sub.Offset = newMsg.UpdateID
		}
		Listner[sub.Id] = sub
	}

}
