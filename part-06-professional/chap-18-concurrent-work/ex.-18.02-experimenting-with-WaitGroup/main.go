package main

import (
	"log"
	"sync"
)

func sum(fromNum, toNum int, wg *sync.WaitGroup, result *int) {
	*result = 0
	for i := fromNum; i <= toNum; i++ {
		*result += i
	}

	wg.Done()

	return
}

func main() {
	batch1 := 0
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go sum(1, 100, wg, &batch1)
	wg.Wait()

	log.Println(batch1)
}
