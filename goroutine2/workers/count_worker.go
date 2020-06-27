package workers

import "fmt"

// CountWorker is ..
type CountWorker struct {
	ID int
}

// Perform is ..
func (w *CountWorker) Perform(c chan int) {
	for {
		data := <-c
		fmt.Printf("Worker %d got data %d\n", w.ID, data)
	}
}
