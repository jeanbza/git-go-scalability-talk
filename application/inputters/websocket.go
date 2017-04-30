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
    port int
}

var upgrader = websocket.Upgrader{}

func NewWebsocketListener(port int) *WebsocketListener {
    return &WebsocketListener{
        port: port,
    }
}

func (l *WebsocketListener) StartAccepting(q queues.Queue) {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        c, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Printf("Upgrade error: %v\n", err)
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
    log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", l.port), nil))
}


