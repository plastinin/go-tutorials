package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var coal atomic.Int64
	var mails []string

	mtx := sync.Mutex{}

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("--->>> Рабочий день шахтеров закончен")
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("--->>> Рабочий день почтальонов закончен")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 300)
	mailTransferPoint := postman.PostmanPool(postmanContext, 300)

	initTime := time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("Угля добыто:", coal.Load())

	mtx.Lock()
	fmt.Println("Писем получено:", len(mails))
	mtx.Unlock()

	fmt.Println("Затраченное время:", time.Since(initTime))
}
