package main

import (
	"github.com/jadekler/git-go-scalability-talk/application/inputters"
	"github.com/jadekler/git-go-scalability-talk/application/outputters"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"sync"
)

type Queue interface {
	Enqueue(data []byte)
	Dequeue() ([]byte, bool)
}

type inputter interface {
	StartAccepting(q queues.Queue)
}

type outputter interface {
	StartOutputting(q queues.Queue)
}

type Processor struct {
	i  inputter
	q  queues.Queue
	o  outputter
	wg *sync.WaitGroup
}

func (p *Processor) Start() {
	go p.i.StartAccepting(p.q)
	go p.o.StartOutputting(p.q)
	p.wg.Wait()
}

func main() {
	//i := listeners.NewWebsocketListener(8080)
	//i := listeners.NewUdpListener(8080)
	//i := listeners.NewStreamingGrpcListener(8080)
	//i := listeners.NewUnaryGrpcListener(8080)
	i := listeners.NewHttpListener(8080)

	//q := queues.NewChannelQueue()
	//q := queues.NewMutexArrayQueue()
	//q := queues.NewMutexRingBufferQueue(10)
	q := queues.NewAtomicRingBuffer(10)

	o := &outputters.StdoutOutputter{}

	p := NewProcessor(i, q, o)
	p.Start()
}

func NewProcessor(i inputter, q queues.Queue, o outputter) *Processor {
	var wg sync.WaitGroup
	wg.Add(1)
	return &Processor{i: i, q: q, o: o, wg: &wg}
}
