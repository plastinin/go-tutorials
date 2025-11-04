package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 5, 7, 4, 3, 2, 1}

	var uniq = 0
	for _, v := range arr {
		uniq ^= v
	}

	fmt.Println(uniq)
}