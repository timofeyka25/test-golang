package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(string(p))
		var msg map[string]string
		err = json.Unmarshal(p, &msg)
		if err != nil {
			log.Println(err)
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
			return
		}

		if msg["event"] == "ping" {
			msg["event"] = "pong"
		}
		responseMsg, err := json.Marshal(msg)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(responseMsg))

		if err := conn.WriteMessage(messageType, responseMsg); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client successfully connected")
	reader(ws)
}

func initRoutes() {
	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	initRoutes()
	log.Fatal(http.ListenAndServe(":28000", nil))
}
