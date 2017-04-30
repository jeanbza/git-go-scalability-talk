package outputters

import (
	"fmt"
	"github.com/jadekler/git-go-scalability-talk/application/queues"
	"time"
)

type StdoutOutputter struct {
}

func (o *StdoutOutputter) StartOutputting(q queues.Queue) {
	fmt.Println("Starting output")
	for {
		for message, ok := q.Dequeue(); ok == true; message, ok = q.Dequeue() {
			fmt.Println(fmt.Sprintf("Got data: %s", string(message)))
		}
		time.Sleep(time.Second)
	}
}
