package main

import (
    "fmt"
    "time"
    "math/rand"
    "sync"
)

var wg sync.WaitGroup

func startProcessor(c chan int) {
    for {
        time.Sleep(time.Duration(rand.Int() % 3) * time.Second)
        fmt.Println(<- c)
        wg.Done()
    }
}

func main() {
    c := make(chan int, 10)

    for i := 0; i < 10; i++ {
        go startProcessor(c)
    }

    for i := 0; i < 100; i++ {
        c <- i
        wg.Add(1)
    }

    fmt.Println("Waiting to finish processing")
    wg.Wait()
    fmt.Println("Done")
}

