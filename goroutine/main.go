package main

import (
	"fmt"
	"time"
)

func mine(transferPoint chan int, n int) {
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")
	transferPoint <- 10
	fmt.Println("Поход номер", n, "уголь передал")
}

func main() {
	coal := 0

	transferPoint := make(chan int, 2)

	before := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)
	
	time.Sleep(3 * time.Second)
	coal += <-transferPoint
	time.Sleep(3 * time.Second)
	coal += <-transferPoint
	time.Sleep(3 * time.Second)
	coal += <-transferPoint

	fmt.Println("Добыли", coal, "угля")
	fmt.Println("Прошло времени", time.Since(before))
}
