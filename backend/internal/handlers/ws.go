package handlers

import (
	"net/http"

	"crypto-tracker/internal/ws"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWS(hub *ws.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := &ws.Client{
		Conn: conn,
		Send: make(chan map[string]float64),
	}

	hub.Register <- client

	go writePump(client)
	go readPump(hub, client)
}

func readPump(hub *ws.Hub, client *ws.Client) {
	defer func() {
		hub.Unregister <- client
		client.Conn.Close()
	}()

	for {
		_, _, err := client.Conn.ReadMessage()
		if err != nil {
			break
		}
	}
}

func writePump(client *ws.Client) {
	for msg := range client.Send {
		client.Conn.WriteJSON(msg)
	}
}

