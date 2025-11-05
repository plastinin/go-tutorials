package main

import (
	"paymentfabric/payments"
	"paymentfabric/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewBank();

	paymentsModule := payments.NewPaymentModule(method)
	paymentsModule.Pay("Бургер", 5)
	idPhone := paymentsModule.Pay("Телефон", 100)
	idGame 	:= paymentsModule.Pay("Игра", 60)

	paymentsModule.Cancel(idPhone)

	allInfo := paymentsModule.AllInfo()

	pp.Println(allInfo)
	pp.Println("Game info:", paymentsModule.Info(idGame))

}