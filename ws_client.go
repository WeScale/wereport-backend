package main

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var listSocketClient []*websocket.Conn

func ClientsWebsocket(ws *websocket.Conn) {
	listSocketClient = append(listSocketClient, ws)
	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)
	}
}

func ClientWebSocketSend(client Client) {
	b, err := json.Marshal(&client)
	if err != nil {
		log.Println("Can't convert")
	}

	for _, socket := range listSocketClient {
		if err = websocket.Message.Send(socket, string(b)); err != nil {
			log.Println("Can't send", err)
			break
		}
	}
}
