package benchmark

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"testing"
)

func BenchmarkStreamingWebsocketListener(b *testing.B) {
	for i := 0; i < b.N; i++ {
		w.wg.Add(1)
		sendPacket(w.c)
	}

	w.wg.Wait()
}

func BenchmarkStreamingWebsocketListenerParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		u := url.URL{Scheme: "ws", Host: fmt.Sprintf("localhost:%d", w.p), Path: "/"}
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			panic(err)
		}

		for pb.Next() {
			w.wg.Add(1)
			sendPacket(c)
		}
	})

	w.wg.Wait()
}

func sendPacket(c *websocket.Conn) {
	err := c.WriteMessage(websocket.TextMessage, []byte(msg))
	if err != nil {
		panic(err)
	}
}
