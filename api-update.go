package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mrzkhd/tlgrm-bot/domain"
	"io/ioutil"
	"net/http"
	"reflect"
)

const (
	TELEGRAM_BASE_URL      = "https://api.telegram.org/SSSSS"
	TELEGRAM_URL_KICK_USER = TELEGRAM_BASE_URL + "kickChatMember"
	APPLICATION_TYPE_JSON  = "application/json"
	CHAT_ID                = "chat_id"
	USER_ID                = "user_id"
)

var update model.Update

func Handler(w http.ResponseWriter, r *http.Request) {
	// Read body
	fmt.Println("Update API is called!")

	b, err := ioutil.ReadAll(r.Body)
	fmt.Println("request body: " + string(b))

	err = json.Unmarshal(b, &update)
	//r.Body.BindJSON(&update)
	if err != nil {
		fmt.Println("err in unmarshaling! " + err.Error())
		http.Error(w, err.Error(), 500)
		return
	}

	if update.Message.From.IsBot == true {
		fmt.Println("Bot detected!")
		kickUser()
	}

	fmt.Fprintf(w, fmt.Sprintf("%d", update.UpdateID))
	//fmt.Println(update.Message.MessageID)

}

func kickUser() {
	defer handlerRecover()

	jsonData := map[string]interface{}{CHAT_ID: update.Message.Chat.ID, USER_ID: update.Message.From.ID}

	jsonValue, _ := json.Marshal(jsonData)
	fmt.Printf("kick user request data: ", jsonValue)

	response, err := http.Post(TELEGRAM_URL_KICK_USER, APPLICATION_TYPE_JSON, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	fmt.Printf("User is kicked", response)

}

func handlerRecover() {
	r := recover()

	if r != nil {

		errorType := reflect.TypeOf(r)

		fmt.Printf("Unexpected error", errorType)

	}
}