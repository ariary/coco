package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Agent struct {
	MachineName string
	Modules     []string
	Websocket   *websocket.Conn
}

type CocoServer struct {
	Port   string
	Agents []Agent
}

var upgrader = websocket.Upgrader{} // use default options

func AgentWebsocket(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("agent connected to coco!"))
	if err != nil {
		log.Println(err)
	}

	Reader(ws) // listen indefinitely for new messages coming through on our bebSocket connection
}

// reader: read message from websocket
func Reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		msg := string(p)
		log.Println("message from agent:", msg)

		// HAndle message
		conn.WriteMessage(messageType, []byte("Thank's"))
	}
}
