package main

import "fmt"

func main() {
	slice := []string{"1", "2", "3"}

	fmt.Println("before:", slice)
	UpdateSlice(slice, "4")
	fmt.Println("after:", slice)
	GrowSlice(slice, "5")
	fmt.Println("afrer:", slice)
}

func UpdateSlice(slice []string, str string) {
	slice[len(slice)-1] = str
	fmt.Println("UpdateSlice:", slice)
}

func GrowSlice(slice []string, str string) {
	slice = append(slice, str)
	fmt.Println("GrowSlice:", slice)
}
