package messaging

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
)

var (
	botToken    = os.Getenv("TELEGRAM_BOT_TOKEN")
	chatID      = os.Getenv("TELEGRAM_CHAT_ID")
	username    = os.Getenv("TELEGRAM_Username")
	channelName = os.Getenv("TELEGRAM_Channel_Name")
)

var int64ChatID, _ = strconv.ParseInt(chatID, 10, 64)

//Base64 encoder
func Encodebase64(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(buff), nil
}

//Test get bot details with the valid data
var _ = Describe("Test get bot details with the valid data", func() {

	os.Setenv("BOT_TOKEN", botToken)
	req, err := http.NewRequest("GET", "/getbot", nil)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBotDetails)
	handler.ServeHTTP(recorder, req)

	Describe("Get Bot Details Telegram", func() {
		Context("get bot details", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})

})

//Test get bot details with the invalid token
var _ = Describe("Test get bot details with the invalid token", func() {
	os.Setenv("BOT_TOKEN", "123")
	req, err := http.NewRequest("GET", "/getbot", nil)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBotDetails)
	handler.ServeHTTP(recorder, req)

	Describe("Get Bot Details Telegram", func() {
		Context("get bot details", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test bot send message
var _ = Describe("Test send bot message with valid data", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{ChatID: int64ChatID, Message: "Test bot send message"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, req)

	Describe("Send bot message", func() {
		Context("send bot message", func() {
			if recorder.Code == http.StatusOK {
				It("Should result http.StatusOK", func() {
					Expect(http.StatusOK).To(Equal(recorder.Code))
				})
			} else {
				It("Should result http.StatusBadRequest", func() {
					Expect(http.StatusBadRequest).To(Equal(recorder.Code))
					fmt.Println("Bot has been removed from group")
				})
			}

		})
	})

})

//Test bot send message invalid chatid
var _ = Describe("Test send bot message with invalid chatid", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{ChatID: -1234, Message: "Test bot send message invalid chatid"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, req)

	Describe("Send bot message", func() {
		Context("send bot message", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test bot send message invalid message type
var _ = Describe("Test send bot message with empty message", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, req)

	Describe("Send bot message", func() {
		Context("send bot message", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test bot send message
var _ = Describe("Test send bot message with Invalid Token", func() {
	botToken := "754194684:A"
	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{ChatID: int64ChatID, Message: "Test bot send message"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendMessage)
	handler.ServeHTTP(recorder, req)

	Describe("Send bot message", func() {
		Context("send bot message", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test bot send message to channel
var _ = Describe("Test send bot channel message with valid data", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{Username: username, Message: "Test send bot channel message"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/sendchannelmessage", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendChannelMessage)
	handler.ServeHTTP(recorder, req)

	Describe("Send bot channel message", func() {
		Context("send bot channel message", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})

})

//Test bot send channel message
var _ = Describe("Test send bot channel message with Invalid username", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{ChatID: 123, Message: "Test send bot channel message"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/sendchannelmessage", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendChannelMessage)
	handler.ServeHTTP(recorder, req)

	Describe("Send bot channel message", func() {
		Context("send bot channel message", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test bot send channel message
var _ = Describe("Test send bot channel message with Invalid data", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/sendchannelmessage", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendChannelMessage)
	handler.ServeHTTP(recorder, req)

	Describe("Send bot channel message", func() {
		Context("send bot channel message", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test bot send message to channel
var _ = Describe("Test send bot channel message with Invalid token", func() {
	botToken := "754194684:A"
	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{Username: username, Message: "Test send bot channel message"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/sendchannelmessage", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendChannelMessage)
	handler.ServeHTTP(recorder, req)

	Describe("Send bot channel message", func() {
		Context("send bot channel message", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test GetChat
var _ = Describe("Test GetChat with valid data", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{ChatID: int64ChatID}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/getchat", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetChat)
	handler.ServeHTTP(recorder, req)

	Describe("Get chat", func() {
		Context("Get chat", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})

})

//Test GetChat
var _ = Describe("Test GetChat with Invalid Chat_Id", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{ChatID: -3492802022}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/getchat", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetChat)
	handler.ServeHTTP(recorder, req)

	Describe("Get chat", func() {
		Context("Get chat", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test GetChat
var _ = Describe("Test GetChat with Invalid data", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/getchat", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetChat)
	handler.ServeHTTP(recorder, req)

	Describe("Get chat", func() {
		Context("Get chat", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test GetChat
var _ = Describe("Test GetChat with Invalid token", func() {
	botToken := "754194684:AA"
	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{Username: "-349280204"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/getchat", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetChat)
	handler.ServeHTTP(recorder, req)

	Describe("Get chat", func() {
		Context("Get chat", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test Send Photo
var _ = Describe("Test send photo with valid data", func() {

	os.Setenv("BOT_TOKEN", botToken)

	filepath, _ := filepath.Abs("../testImage/dice.jpeg")
	base64Data, base64Err := Encodebase64(filepath)
	if base64Err != nil {
		fmt.Println("===base64 err======", base64Err)
	}

	botMessage := BotMessage{ChatID: int64ChatID, ImageBase64: base64Data}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/sendphoto", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendPhoto)
	handler.ServeHTTP(recorder, req)

	Describe("Test Send Photo", func() {
		Context("Send Photo", func() {
			if recorder.Code == http.StatusOK {
				It("Should result http.StatusOK", func() {
					Expect(http.StatusOK).To(Equal(recorder.Code))
				})
			} else {
				It("Should result http.StatusBadRequest", func() {
					Expect(http.StatusBadRequest).To(Equal(recorder.Code))
					fmt.Println("Photo not sent!! Bot left the group")
				})
			}
		})
	})

})

//Test Send Photo
var _ = Describe("Test send photo with Invalid Chat_Id", func() {

	os.Setenv("BOT_TOKEN", botToken)

	filepath, _ := filepath.Abs("../testImage/dice.jpeg")
	base64Data, base64Err := Encodebase64(filepath)
	if base64Err != nil {
		fmt.Println("===base64 err======", base64Err)
	}

	botMessage := BotMessage{ChatID: -3492802, ImageBase64: base64Data}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/sendphoto", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendPhoto)
	handler.ServeHTTP(recorder, req)

	Describe("Test Send Photo", func() {
		Context("Send Photo", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test Send Photo
var _ = Describe("Test send photo with Invalid data", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/sendphoto", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendPhoto)
	handler.ServeHTTP(recorder, req)

	Describe("Test Send Photo", func() {
		Context("Send Photo", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test Send Photo
var _ = Describe("Test send photo with Invalid token", func() {
	botToken := "754194684:AA"
	os.Setenv("BOT_TOKEN", botToken)

	filepath, _ := filepath.Abs("../testImage/dice.jpeg")
	base64Data, base64Err := Encodebase64(filepath)
	if base64Err != nil {
		fmt.Println("===base64 err======", base64Err)
	}

	botMessage := BotMessage{ChatID: int64ChatID, ImageBase64: base64Data}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/sendphoto", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendPhoto)
	handler.ServeHTTP(recorder, req)

	Describe("Test Send Photo", func() {
		Context("Send Photo", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Subscribe
var _ = Describe("Subscribe Updates", func() {

	os.Setenv("BOT_TOKEN", botToken)
	data := Data{
		Channel: channelName,
	}
	botMessage := Subscribe{Endpoint: "https://webhook.site/3cee781d-0a87-4966-bdec-9635436294e9",
		Id:        "1",
		IsTesting: true,
		Data:      data,
	}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		fmt.Println(" request err :", err)
	}
	req, err := http.NewRequest("POST", "/subscribe", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SubscribeUpdate)
	handler.ServeHTTP(recorder, req)

	Describe("Subscribe", func() {
		Context("Subscribe", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

//Test Leave Chat
var _ = Describe("Test Leave Chat with valid data", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{ChatID: int64ChatID}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/leavechat", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(LeaveChat)
	handler.ServeHTTP(recorder, req)

	Describe("LeaveChat", func() {
		Context("Leave Chat", func() {
			if recorder.Code == http.StatusOK {
				It("Should result http.StatusOK", func() {
					Expect(http.StatusOK).To(Equal(recorder.Code))
				})
			} else {
				It("Should result http.StatusBadRequest", func() {
					Expect(http.StatusBadRequest).To(Equal(recorder.Code))
				})
			}

		})
	})

})

//Test Leave Chat
var _ = Describe("Test Leave Chat with Invalid Chat_Id", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{ChatID: -349}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/leavechat", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(LeaveChat)
	handler.ServeHTTP(recorder, req)

	Describe("LeaveChat", func() {
		Context("Leave Chat", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})

		})
	})

})

//Test Leave Chat
var _ = Describe("Test Leave Chat with Invalid data", func() {

	os.Setenv("BOT_TOKEN", botToken)
	botMessage := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/leavechat", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(LeaveChat)
	handler.ServeHTTP(recorder, req)

	Describe("LeaveChat", func() {
		Context("Leave Chat", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test Leave Chat
var _ = Describe("Test Leave Chat with Invalid token", func() {
	botToken := "754194684:A"
	os.Setenv("BOT_TOKEN", botToken)
	botMessage := BotMessage{Message: "-349280204"}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(botMessage)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/leavechat", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(LeaveChat)
	handler.ServeHTTP(recorder, req)

	Describe("LeaveChat", func() {
		Context("Leave Chat", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})
