package main

import (
    "log"
    "net/url"
    "github.com/gorilla/websocket"
)

// Sends a single message over websocket
func main() {
    u := url.URL{Scheme: "ws", Host: "localhost:8000", Path: "/echo"}
    log.Printf("connecting to %s", u.String())

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
