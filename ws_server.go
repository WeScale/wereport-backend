package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func WebSocketStart() {
	http.Handle("/", websocket.Handler(Echo))
	http.Handle("/consultants", websocket.Handler(ConsultantsWebsocket))
	http.Handle("/clients", websocket.Handler(ClientsWebsocket))

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
