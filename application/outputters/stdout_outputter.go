package outputters

import (
    "github.com/jadekler/git-go-scalability-talk/application/queues"
    "fmt"
)

type StdoutOutputter struct {

}

func (o *StdoutOutputter) StartOutputting(q queues.Queue) {
    fmt.Println("Starting output")
    for {
        message := q.Dequeue()
        fmt.Println(string(message))
    }
}
