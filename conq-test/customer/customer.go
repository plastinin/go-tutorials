package customer

import (
	"conq/common"
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func customer(n int, demand chan<- common.Operation, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("done: Покупатель %d завершил работу\n", n)
			return
		default:
			amount := rand.IntN(1000) + 1
			response := make(chan bool)
			demand <- common.Operation{Amount: amount, Responce: response}

			// Ждем ответа от склада
			if <-response {
				time.Sleep(time.Duration(rand.IntN(5)) * time.Second)
			}
		}
	}
}

func CustomerPool(count int, ctx context.Context) <-chan common.Operation {

	demandChan := make(chan common.Operation, 10)

	wg := &sync.WaitGroup{}
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go customer(i, demandChan, wg, ctx)
	}

	go func() {
		wg.Wait()
		close(demandChan)
	}()

	return demandChan
}
