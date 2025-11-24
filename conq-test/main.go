package main

import (
	"conq/customer"
	"conq/partner"
	"conq/warehouse"
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {

	partnerContext, partnerClose := context.WithCancel(context.Background())

	go func() {
		time.Sleep(30 * time.Second)
		fmt.Println("--->>> Поставки завершены")
		partnerClose()
	}()

	supplyChan := partner.PartnerPool(rand.IntN(500_000)+1, partnerContext)
	demandChan := customer.CustomerPool(rand.IntN(1_000_000)+1, partnerContext)

	warehouse.WarehouseManager(supplyChan, demandChan)
}
