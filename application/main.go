package main

import (
    "github.com/jadekler/git-go-scalability-talk/application/queues"
    "github.com/jadekler/git-go-scalability-talk/application/inputters"
    "github.com/jadekler/git-go-scalability-talk/application/outputters"
)

type inputter interface {
    StartAccepting(q queues.Queue)
}

type outputter interface {
    StartOutputting(q queues.Queue)
}

func main() {
    i := listeners.NewWebsocketListener("localhost:8080")
    q := queues.NewChanneler()
    o := &outputters.StdoutOutputter{}
    p := NewProcessor(i, q, o)
    p.Start()
}

type Processor struct {
    i inputter
    q queues.Queue
    o outputter
}

func NewProcessor(i inputter, q queues.Queue, o outputter) *Processor {
    return &Processor{i: i, q: q, o: o}
}

func (p *Processor) Start() {
    go p.i.StartAccepting(p.q)
    go p.o.StartOutputting(p.q)
}