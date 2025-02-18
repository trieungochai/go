// In this exercise, we will calculate a sum using several workers.
// A worker is essentially a function,
// and we will be organizing these workers into a single struct
package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	in, out chan int
	sbw     int // sbw: subworker
	mtx     *sync.Mutex
}

// Create a method and increment the number of subworker instances.
// Sub-workers are basically identical Goroutines that split the work that needs to be done.
// Note that the function is meant to be used directly and not as a Goroutine,
// as it itself creates a new Goroutine.
func (w *Worker) readThem() {
	w.sbw++
	go func() {
		partial := 0
		for i := range w.in {
			partial += i
		}
		w.out <- partial

		w.mtx.Lock()
		w.sbw--
		if w.sbw == 0 {
			close(w.out)
		}
		w.mtx.Unlock()
	}()

}

func (w *Worker) gatherResult() int {
	total := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := range w.out {
			total += i
		}
		wg.Done()
	}()

	wg.Wait()
	return total
}

func main() {

	mtx := &sync.Mutex{}
	in := make(chan int, 100)
	wrNum := 10
	out := make(chan int)
	wrk := Worker{in: in, out: out, mtx: mtx}

	for i := 1; i <= wrNum; i++ {
		wrk.readThem()
	}

	for i := 1; i <= 100; i++ {
		in <- i
	}
	close(in)

	res := wrk.gatherResult()

	fmt.Println(res)
}
