package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var number atomic.Int64

func increase(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 1_000; i++ {
		number.Add(1)
	}
}

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(10)

	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)
	go increase(wg)

	wg.Wait()
	fmt.Println("number:", number.Load())
}
