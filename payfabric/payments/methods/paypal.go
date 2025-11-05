package methods

import (
	"fmt"
	"math/rand/v2"
)

type Paypal struct {
}

func NewPayPal() Paypal {
	return Paypal{}
}

func (c Paypal) Pay(usd int) int {
	fmt.Println("Оплата PayPal")
	fmt.Println("Размер оплаты", usd, "USD")

	return rand.Int()
}

func (c Paypal) Cancel(id int) {
	fmt.Println("Отмена PayPal! ID:", id)
}