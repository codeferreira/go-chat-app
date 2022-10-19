package main

import (
	"fmt"
	"net/http"

	"github.com/codeferreira/realtime-chat-app/pkg/websocket"
)

func serverWs(pool *websocket.Pool, writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Host)
	connection, err := websocket.Upgrade(writer, request)

	if err != nil {
		fmt.Println(writer, "%+V\n", err)
	}

	client := &websocket.Client{
		Conn: connection,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	fs := http.FileServer(http.Dir("./web"))

	http.Handle("/", fs)

	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		serverWs(pool, writer, request)
	})
}

func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":3333", nil)
}
