package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

// Sends a single message over websocket
func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:8000", Path: "/echo"}
	log.Printf("Connecting to %s\n", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "Hello world!"))
	if err != nil {
		log.Println("write close:", err)
		return
	}
}
