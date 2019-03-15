package messaging

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//Test get bot details with the valid data
var _ = Describe("Test get bot details with the valid data", func() {
	accessToken := "754194684:AAESS4D5lHbhOW8Gs4eBiO3ZNSfaCYl1tMA"
	os.Setenv("ACCESS_TOKEN", accessToken)
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
	os.Setenv("ACCESS_TOKEN", "123")
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
	accessToken := "754194684:AAESS4D5lHbhOW8Gs4eBiO3ZNSfaCYl1tMA"
	os.Setenv("ACCESS_TOKEN", accessToken)
	botMessage := BotMessage{ChatID: -349280204, Message: "Test bot send message"}
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(botMessage)
	req, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, req)

	Describe("Send bot message", func() {
		Context("send bot message", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})

})

//Test bot send message invalid chatid
var _ = Describe("Test send bot message with invalid chatid", func() {
	accessToken := "754194684:AAESS4D5lHbhOW8Gs4eBiO3ZNSfaCYl1tMA"
	os.Setenv("ACCESS_TOKEN", accessToken)
	botMessage := BotMessage{ChatID: -1234, Message: "Test bot send message"}
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(botMessage)
	req, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, req)

	Describe("Send bot message", func() {
		Context("send bot message", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})

//Test bot send message invalid message
var _ = Describe("Test send bot message with invalid message", func() {
	accessToken := "754194684:AAESS4D5lHbhOW8Gs4eBiO3ZNSfaCYl1tMA"
	os.Setenv("ACCESS_TOKEN", accessToken)
	botMessage := BotMessage{ChatID: -349280204, Message: ""}
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(botMessage)
	req, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
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
var _ = Describe("Test send bot message with valid data", func() {
	accessToken := "754194684:AAESS4D5lHbhOW8Gs4eBiO3ZNSfaCYl1tMAA"
	os.Setenv("ACCESS_TOKEN", accessToken)
	botMessage := BotMessage{ChatID: -349280204, Message: "Test bot send message"}
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(botMessage)
	req, err := http.NewRequest("POST", "/send", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Send)
	handler.ServeHTTP(recorder, req)

	Describe("Send bot message", func() {
		Context("send bot message", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})

})
