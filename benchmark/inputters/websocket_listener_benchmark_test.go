package benchmark

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jadekler/git-go-scalability-talk/application/inputters"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"net/url"
	"sync"
	"testing"
)

var w websocketListenerBenchmark = websocketListenerBenchmark{}

type websocketListenerBenchmark struct {
	l  *listeners.WebsocketListener
	wg *sync.WaitGroup
	q  queues.Queue
	p  int
    c *websocket.Conn
}

func BenchmarkWebsocketListener(b *testing.B) {
	if w.l == nil {
		w.p = benchmark.GetOpenTcpPort()
		fmt.Println("Starting on port", w.p)

		w.wg = &sync.WaitGroup{}
		w.q = benchmark.NewWaitingQueue(w.wg)

		w.l = listeners.NewWebsocketListener(w.p)
		go w.l.StartAccepting(w.q)

        u := url.URL{Scheme: "ws", Host: fmt.Sprintf("localhost:%d", w.p), Path: "/"}
        fmt.Println("Connecting to", u.String())
        var err error
        w.c, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
        if err != nil {
            panic(err)
        }
	}

	for i := 0; i < b.N; i++ {
		w.wg.Add(1)
        sendPacket(w.p, w.c)
	}

	w.wg.Wait()
}

func sendPacket(port int, c *websocket.Conn) {
    err := c.WriteMessage(websocket.TextMessage, []byte(LARGE_MESSAGE))
    if err != nil {
        panic(err)
    }
}
