package main

import (
	"encoding/json"
	"fmt"
	"log"

	"golang.org/x/net/websocket"
)

var listSocketContrats []*websocket.Conn

func ContratsWebsocket(ws *websocket.Conn) {
	log.Println("add client for clientsocket")
	listSocketContrats = append(listSocketContrats, ws)
	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			break
		}

		fmt.Println("Received back from client: " + reply)
	}
}

func ContratWebSocketSend(contrat Contrat) {
	b, err := json.Marshal(&contrat)
	if err != nil {
		log.Println("Can't convert")
	}

	for i, socket := range listSocketContrats {
		log.Println("send contrat for contratsocket", string(b))
		if err = websocket.Message.Send(socket, string(b)); err != nil {
			log.Println("Can't send", err)
			listSocketContrats = append(listSocketContrats[:i], listSocketContrats[i+1:]...)
		}
	}
}
