package main

import (
	"fmt"
	"net/http"
)


func HandlerHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1> Server is running </h1>")
	fmt.Fprintf(w, "<h2> Telegram bot is Active. </h2>")

}