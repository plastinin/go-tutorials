package main

import (
	"cmp"
	"fmt"
)

func main() {

	var x, y int

	x = 2
	y = 6

	fmt.Println(Max(x, y))

	list := []int{1, 2, 4, 8}
	fmt.Println(IsContains(1, list))
	fmt.Println(IsContains(5, list))

	list2 := []string{"apple", "banana", "orange"}
	fmt.Println(IsContains("apple", list2))
	fmt.Println(IsContains("mango", list2))

	fmt.Println("sum:", Sum(list))

	sum := func(k1, k2 int) int { return k1 + k2}
	fmt.Println("reduce:", Reduce(list, sum, 0))

}

// type Ordered interface {
// 	int | uint
// }

func Max[T cmp.Ordered](x, y T) T {
	return Ternary(x > y, x, y)
}

func IsContains[T comparable](n T, list []T) bool {
	for _, v := range list {
		if v == n {
			return true
		}
	}

	return false
}

func Sum[T cmp.Ordered](list[]T) T {
	var res T

	for _, value := range list {
		res += value
	}

	return res
}

func Reduce[T any](list[]T, accumulator func(T, T) T, init T) T {
	for _, v := range list {
		init = accumulator(init, v)
	}
	return init
}

func Ternary[T any](cond bool, x T, y T) T {
	if cond {
		return x 
	}

	return y
}
