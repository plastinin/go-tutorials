package main

import (
	"fmt"
	"time"
)

func mine(n int) int {
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(10 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")
	return 10
}

func main() {
	coal := 0

	before := time.Now()

	coal += mine(1)
	coal += mine(2)
	coal += mine(3)

	fmt.Println("Добыли", coal, "угля")
	fmt.Println("Прошло времени", time.Since(before))
}
