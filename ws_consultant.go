package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

var listSocketConsultant []*websocket.Conn

func ConsultantsWebsocket(ws *websocket.Conn) {
	listSocketConsultant = append(listSocketConsultant, ws)
	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)
	}
}

func ConsultantWebSocketSend(consultant Consultant) {
	b, err := json.Marshal(&consultant)
	if err != nil {
		log.Println("Can't convert")
	}

	for _, socket := range listSocketConsultant {
		if err = websocket.Message.Send(socket, string(b)); err != nil {
			log.Println("Can't send", err)
			break
		}
	}
}
