package benchmark

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jadekler/git-go-scalability-talk/application/inputters"
	"github.com/jadekler/git-go-scalability-talk/application/model"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"github.com/jadekler/git-go-scalability-talk/benchmark"
	"google.golang.org/grpc"
	"net"
	"net/url"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

var h httpListenerBenchmark = httpListenerBenchmark{}

type httpListenerBenchmark struct {
	l  *listeners.HttpListener
	wg *sync.WaitGroup
	q  queues.Queue
	p  int
}

var t tcpListenerBenchmark = tcpListenerBenchmark{}

type tcpListenerBenchmark struct {
	l    *listeners.TcpListener
	wg   *sync.WaitGroup
	q    queues.Queue
	p    int
	conn net.Conn
}

var sg streamingGrpcListenerBenchmark = streamingGrpcListenerBenchmark{}

type streamingGrpcListenerBenchmark struct {
	l  *listeners.StreamingGrpcListener
	wg *sync.WaitGroup
	q  queues.Queue
	p  int
	s  model.GrpcStreamingInputterService_MakeRequestClient
}

var su streamingUdpListenerBenchmark = streamingUdpListenerBenchmark{}

type streamingUdpListenerBenchmark struct {
	l    *listeners.UdpListener
	wg   *sync.WaitGroup
	t    *time.Timer
	q    queues.Queue
	p    int
	conn *net.UDPConn
}

var uu unaryUdpListenerBenchmark = unaryUdpListenerBenchmark{}

type unaryUdpListenerBenchmark struct {
	l     *listeners.UdpListener
	wg    *sync.WaitGroup
	t     *time.Timer
	q     queues.Queue
	p     int
	laddr *net.UDPAddr
	raddr *net.UDPAddr
}

var ug unaryGrpcListenerBenchmark = unaryGrpcListenerBenchmark{}

type unaryGrpcListenerBenchmark struct {
	l  *listeners.UnaryGrpcListener
	wg *sync.WaitGroup
	q  queues.Queue
	p  int
	c  model.GrpcUnaryInputterServiceClient
}

var w websocketListenerBenchmark = websocketListenerBenchmark{}

type websocketListenerBenchmark struct {
	l  *listeners.WebsocketListener
	wg *sync.WaitGroup
	q  queues.Queue
	p  int
	c  *websocket.Conn
}

func TestMain(m *testing.M) {
	fmt.Println("Setup!")

	setupHttp()
	setupTcp()
	setupStreamingGrpc()
	setupStreamingUdp()
	setupUnaryUdp()
	setupUnaryGrpc()
	setupWebsocket()

	os.Exit(m.Run())
}

func setupTcp() {
	t.p = benchmark.GetOpenTcpPort()
	t.wg = &sync.WaitGroup{}
	t.q = benchmark.NewWaitingQueue(t.wg)
	t.l = listeners.NewTcpListener(t.p)
	go t.l.StartAccepting(t.q)
	t.conn = openTcpConn(t.p)
}

func setupHttp() {
	h.p = benchmark.GetOpenTcpPort()
	h.wg = &sync.WaitGroup{}
	h.q = benchmark.NewWaitingQueue(h.wg)
	h.l = listeners.NewHttpListener(h.p)
	go h.l.StartAccepting(h.q)
}

func setupStreamingGrpc() {
	sg.p = benchmark.GetOpenTcpPort()

	sg.wg = &sync.WaitGroup{}
	sg.q = benchmark.NewWaitingQueue(sg.wg)

	sg.l = listeners.NewStreamingGrpcListener(sg.p)
	go sg.l.StartAccepting(sg.q)

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", sg.p), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := model.NewGrpcStreamingInputterServiceClient(conn)
	stream, err := client.MakeRequest(context.Background())
	if err != nil {
		panic(err)
	}
	sg.s = stream
}

func setupStreamingUdp() {
	su.p = benchmark.GetOpenTcpPort()

	timerDuration := 50 * time.Millisecond
	su.t = time.NewTimer(timerDuration) // max time we'll wait for a udp packet

	su.wg = &sync.WaitGroup{}
	su.q = benchmark.NewTimeoutQueue(su.wg, su.t, &timerDuration)

	su.l = listeners.NewUdpListener(su.p)
	go su.l.StartAccepting(su.q)

	raddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", su.p))
	if err != nil {
		panic(err)
	}

	laddr, err := net.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		panic(err)
	}

	su.conn = conn
}

func setupUnaryUdp() {
	uu.p = benchmark.GetOpenTcpPort()

	timerDuration := 50 * time.Millisecond
	uu.t = time.NewTimer(timerDuration) // max time we'll wait for a udp packet

	uu.wg = &sync.WaitGroup{}
	uu.q = benchmark.NewTimeoutQueue(uu.wg, uu.t, &timerDuration)

	uu.l = listeners.NewUdpListener(uu.p)
	go uu.l.StartAccepting(uu.q)

	raddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%d", uu.p))
	if err != nil {
		panic(err)
	}

	laddr, err := net.ResolveUDPAddr("udp", "localhost:0")
	if err != nil {
		panic(err)
	}

	uu.raddr = raddr
	uu.laddr = laddr
}

func setupUnaryGrpc() {
	ug.p = benchmark.GetOpenTcpPort()

	ug.wg = &sync.WaitGroup{}
	ug.q = benchmark.NewWaitingQueue(ug.wg)

	ug.l = listeners.NewUnaryGrpcListener(ug.p)
	go ug.l.StartAccepting(ug.q)

	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", ug.p), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	ug.c = model.NewGrpcUnaryInputterServiceClient(conn)
}

func setupWebsocket() {
	w.p = benchmark.GetOpenTcpPort()

	w.wg = &sync.WaitGroup{}
	w.q = benchmark.NewWaitingQueue(w.wg)

	w.l = listeners.NewWebsocketListener(w.p)
	go w.l.StartAccepting(w.q)

	u := url.URL{Scheme: "ws", Host: fmt.Sprintf("localhost:%d", w.p), Path: "/"}
	var err error
	w.c, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}
}

// Open with some minimal retry
func openTcpConn(port int) net.Conn {
	var err error

	for i := 0; i < 5; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", port))

		if err != nil {
			if strings.Contains(err.Error(), "getsockopt: connection refused") {
				fmt.Println("Retrying conn")
				time.Sleep(100 * time.Millisecond)
				continue
			}

			panic(err)
		}

		return conn
	}

	panic(err)
}
