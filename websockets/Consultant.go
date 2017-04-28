package Websockets

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/WeScale/wereport-backend/data"
	"golang.org/x/net/websocket"
)

var listSocketConsultant []*websocket.Conn

func ConsultantsWebsocket(ws *websocket.Conn) {
	log.Println("add client for consultantsocket")
	listSocketConsultant = append(listSocketConsultant, ws)
	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			break
		}

		fmt.Println("Received back from client: " + reply)
	}
}

func ConsultantWebSocketSend(consultant Data.Consultant) {
	b, err := json.Marshal(&consultant)
	if err != nil {
		log.Println("Can't convert", err)
	}

	for i, socket := range listSocketConsultant {
		log.Println("send client for consultantsocket", string(b))
		if err = websocket.Message.Send(socket, string(b)); err != nil {
			log.Println("Can't send", err)
			listSocketConsultant = append(listSocketConsultant[:i], listSocketConsultant[i+1:]...)
		}
	}
}
