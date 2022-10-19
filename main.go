package main

import (
	"fmt"
	"net/http"

	"github.com/codeferreira/realtime-chat-app/pkg/websocket"
)

func serverWs(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Host)

	ws, err := websocket.Upgrade(writer, request)

	if err != nil {
		fmt.Println(writer, "%+V\n", err)
	}

	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Simple Server")
	})

	http.HandleFunc("/ws", serverWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":3333", nil)
}
