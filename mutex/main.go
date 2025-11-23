package main

import (
	"fmt"
	"sync"
)

var slice []int

var mtx sync.Mutex

func increase(wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 1_000; i++ {
		mtx.Lock()
		slice = append(slice, i)
		mtx.Unlock()
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
	fmt.Println("number:", len(slice))
}
