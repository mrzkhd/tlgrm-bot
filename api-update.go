package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mrzkhd/tlgrm-bot/domain"
	"io/ioutil"
	"net/http"
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
	b, err := ioutil.ReadAll(r.Body)

	err = json.Unmarshal(b, &update)
	//r.Body.BindJSON(&update)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	/*if update.Message.From.IsBot == true {
		fmt.Println("Bot detected!")
		kickUser()
	}*/

	fmt.Fprintf(w, string(update.UpdateID))
	//fmt.Println(update.Message.MessageID)

}

func kickUser() {

	jsonData := map[string]interface{}{CHAT_ID: update.Message.Chat.ID, USER_ID: update.Message.From.ID}

	jsonValue, _ := json.Marshal(jsonData)

	response, err := http.Post(TELEGRAM_URL_KICK_USER, APPLICATION_TYPE_JSON, bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	fmt.Printf("Used is kicked", response)

}
