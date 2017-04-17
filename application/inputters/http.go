// Used code from: https://golang.org/pkg/net/
package listeners

import (
    "net/http"
    "fmt"
    "log"
    "github.com/jadekler/git-go-scalability-talk/application/queues"
    "io/ioutil"
)

type HttpListener struct {
    port int
}

func NewHttpListener(port int) *HttpListener {
    return &HttpListener{port: port}
}

func (l *HttpListener) StartAccepting(q queues.Queue) {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
            log.Fatal(err)
        }
        q.Enqueue(body)
    })

    fmt.Println("Accepting input")

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", l.port), nil))
}


