package main

import (
	"goroutine2/workers"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		worker := &workers.CountWorker{
			ID: i,
		}
		go worker.Perform(c)
	}

	for j := 0; j < 5; j++ {
		c <- rand.Int()
		time.Sleep(time.Millisecond * 50)
	}
}
