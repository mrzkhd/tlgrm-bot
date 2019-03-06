package main

import (
	"fmt"
	"net/http"
)

func Handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> Server is running </h1>")

}
