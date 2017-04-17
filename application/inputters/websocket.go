// Used code from: https://github.com/gorilla/websocket/tree/master/examples/echo
package listeners

import (
    "log"
    "net/http"

    "github.com/gorilla/websocket"
    "github.com/jadekler/git-go-scalability-talk/application/queues"
    "fmt"
)

type WebsocketListener struct {
    address string
}

var upgrader = websocket.Upgrader{}

func NewWebsocketListener(address string) *WebsocketListener {
    return &WebsocketListener{
        address: address,
    }
}

func (l *WebsocketListener) StartAccepting(q queues.Queue) {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        c, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Print("Upgrade error:", err)
            return
        }
        log.Print("Upgraded!")
        defer c.Close()
        for {
            _, message, err := c.ReadMessage()
            if err != nil {
                log.Println("read:", err)
                break
            }
            q.Enqueue(message)
        }
    })
    fmt.Println("Accepting input")
    log.Fatal(http.ListenAndServe(l.address, nil))
}


