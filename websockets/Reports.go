package Websockets

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/WeScale/wereport-backend/data"
	"golang.org/x/net/websocket"
)

var listSocketReports []*websocket.Conn

func ReportsWebsocket(ws *websocket.Conn) {
	log.Println("add client for Reportsocket")
	listSocketReports = append(listSocketReports, ws)
	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			break
		}

		fmt.Println("Received back from client: " + reply)
	}
}

func ReportWebSocketSend(report Data.Report) {
	b, err := json.Marshal(&report)
	if err != nil {
		log.Println("Can't convert")
	}

	for i, socket := range listSocketReports {
		log.Println("send reports for Reportsocket", string(b))
		if err = websocket.Message.Send(socket, string(b)); err != nil {
			log.Println("Can't send", err)
			listSocketReports = append(listSocketReports[:i], listSocketReports[i+1:]...)
		}
	}
}
