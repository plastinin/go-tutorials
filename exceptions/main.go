package main

import (
	"errors"
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name     string
	Ballance int
}

func Pay(user *User, usd int) (int, error) {

	if user.Ballance-usd < 0 {
		return user.Ballance, errors.New("Недостаточно средств!")
	}

	user.Ballance -= usd

	return user.Ballance, nil
}

func main() {

	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("Panic!", p)
		}
	}()

	a := 0
	b := 1 / a // panic is here
	fmt.Println(b)

	user := User{
		Name:     "Олег",
		Ballance: 40,
	}
	pp.Println("Before:", user)
	currBalance, err := Pay(&user, 50)
	if err != nil {
		fmt.Println("Оплаты не было! Причина: ", err.Error())
	}

	pp.Println("Currbalance:", currBalance)
}
