package main

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var listSocketClient []*websocket.Conn

func ClientsWebsocket(ws *websocket.Conn) {
	log.Println("add client for clientsocket")
	listSocketClient = append(listSocketClient, ws)
	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			break
		}

		fmt.Println("Received back from client: " + reply)
	}
}

func ClientWebSocketSend(client Client) {
	b, err := json.Marshal(&client)
	if err != nil {
		log.Println("Can't convert", err)
	}

	for i, socket := range listSocketClient {
		log.Println("send client for clientsocket", string(b))
		if err = websocket.Message.Send(socket, string(b)); err != nil {
			log.Println("Can't send", err)
			listSocketClient = append(listSocketClient[:i], listSocketClient[i+1:]...)
		}
	}
}
