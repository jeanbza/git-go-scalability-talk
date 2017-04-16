package queues

type Channeler struct {

}

func NewChanneler() *Channeler {
    return &Channeler{}
}

func (q *Channeler) Enqueue(data []byte) {

}

func (q *Channeler) Dequeue() []byte {
    return nil
}
