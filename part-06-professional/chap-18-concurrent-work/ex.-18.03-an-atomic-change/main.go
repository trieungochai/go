package main

import (
	"log"
	"sync"
	"sync/atomic"
)

// we just changed res from int to *int32.
// The reason for this is that
// the atomic operations available specifically for arithmetic operations only work on int32/64 and relative uint32/64.
func sum(fromNum, toNum int, wg *sync.WaitGroup, result *int32) {
	for i := fromNum; i <= toNum; i++ {
		// instead of assigning the value of res as 0, we are now adding i to the total value held by res.
		// The rest of the code is unchanged.
		atomic.AddInt32(result, int32(i))
	}

	wg.Done()
	return
}

func main() {
	s1 := int32(0)
	wg := &sync.WaitGroup{}
	wg.Add(4)

	go sum(1, 25, wg, &s1)
	go sum(26, 50, wg, &s1)
	go sum(51, 75, wg, &s1)
	go sum(76, 100, wg, &s1)

	wg.Wait()
	log.Println(s1)
}
