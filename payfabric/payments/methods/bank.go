package methods

import (
	"fmt"
	"math/rand/v2"
)

type Bank struct {
}

func NewBank() Bank {
	return Bank{}
}

func (c Bank) Pay(usd int) int {
	fmt.Println("Оплата банковской картой")
	fmt.Println("Размер оплаты", usd, "USD")

	return rand.Int()
}

func (c Bank) Cancel(id int) {
	fmt.Println("Отмена банковской операции! ID:", id)
}