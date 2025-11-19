package main

import (
	"fmt"
	"strconv"
)

type Ordered interface {
	~int | ~float64
}

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintInt int

func (pi PrintInt) String() string {
	return strconv.Itoa(int(pi))
}

type PrintFloat float64

func (pf PrintFloat) String() string {
	return fmt.Sprintf("%f", pf)
}

func main() {

	var i PrintInt = 20
	var f PrintFloat = 10.23

	slice := CreateSlice()

	for _, v := range slice {
		fmt.Println("value:", v)
	}

	fmt.Println("double 2:", Double(2))
	fmt.Println("double 2.5:", Double(2.5))

	Print(i)
	Print(f)

}

func CreateSlice() []any {
	newList := []any{}
	newList = append(newList, 1)
	newList = append(newList, "привет")
	return newList
}

func Double[T Ordered](v T) T {
	return (v * 2)
}

func Print[T Printable](v T) {
	fmt.Println(v)
}
