package websocket

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(writer http.ResponseWriter, request *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(writer, request, nil)

	if err != nil {
		log.Println(err)
		return ws, err
	}

	return ws, nil
}

func Reader(connection *websocket.Conn) {
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

func Writer(connection *websocket.Conn) {
	for {
		fmt.Println("Sending")
		messageType, reader, err := connection.NextReader()

		if err != nil {
			fmt.Println(err)
			return
		}

		w, err := connection.NextWriter(messageType)

		if err != nil {
			fmt.Println(err)
			return
		}

		if _, err := io.Copy(w, reader); err != nil {
			fmt.Println(err)
			return
		}

		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}