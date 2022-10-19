package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func reader(connection *websocket.Conn) {
	for {
		messageType, message, err := connection.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(message))

		if err := connection.WriteMessage(messageType, message); err != nil {
			log.Println(err)
			return
		}
	}
}

func serverWs(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Host)

	ws, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		log.Println(err)
	}

	reader(ws)
}