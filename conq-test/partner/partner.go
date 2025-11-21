package partner

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func partner(n int, supply chan<- int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("🏭 Поставщик %d завершил работу\n", n)
			return
		default:
			supply <- rand.IntN(1000) + 1
			time.Sleep(time.Duration(rand.IntN(5)) * time.Second)
		}
	}
}

func PartnerPool(count int, ctx context.Context) <-chan int {

	supply := make(chan int, 10)

	wg := &sync.WaitGroup{}
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go partner(i, supply, wg, ctx)
	}

	go func() {
		wg.Wait()
		close(supply)
	}()

	return supply
}
