package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(text string, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 3; i++ {
		fmt.Println("Я почтальон, я отнес газету", text, "в", i, "раз")
		time.Sleep(250 * time.Millisecond)
	}

}

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go postman("Новости", wg)

	wg.Add(1)
	go postman("Игровой журнал", wg)

	wg.Add(1)
	go postman("Автомобильная хроника", wg)

	wg.Wait()

	fmt.Println("main завершился")
}
