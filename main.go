package main

import (
	"log"
	"net/http"

	"github.com/WeScale/wereport-backend/websockets"
	"github.com/gorilla/handlers"
)

func main() {

	go Websockets.WebSocketStart()

	router := NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"authorization", "content-type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}
