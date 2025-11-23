package warehouse

import (
	"conq/common"
	"errors"
	"fmt"
)

type Warehouse struct {
	balance int
}

func (wh *Warehouse) Supply(count int) {
	wh.balance += count
}

func (wh *Warehouse) Demand(count int) error {
	if r := wh.balance - count; r < 0 {
		return errors.New("NO_LIMIT")
	} else {
		wh.balance = r
	}
	return nil
}

// supply - очередь на получение
// demand - очередь на списание
func WarehouseManager(supply <-chan int, demand <-chan common.Operation) {

	wh := &Warehouse{balance: 0}
	pending := []common.Operation{} // очередь ожидающих заказов

	for {
		select {
		case amount, ok := <-supply:
			if !ok {
				return
			}
			wh.Supply(amount)
			fmt.Printf("📦 Поставка: +%d, Остаток: %d\n", amount, wh.balance)

			// при поступлении обработаем очередь ожидающих заказов
			pending = processPending(pending, wh)

		case op, ok := <-demand:
			if !ok {
				return
			}
			if err := wh.Demand(op.Amount); err != nil {
				fmt.Printf("⏳ Недостаточно товара для заказа на %d (есть %d)\n",
					op.Amount, wh.balance)
				pending = append(pending, op)
			} else {
				fmt.Printf("🛒 Отгрузка: -%d, Остаток: %d\n", op.Amount, wh.balance)
				op.Responce <- true
			}
		}
	}
}

func processPending(pending []common.Operation, wh *Warehouse) []common.Operation {
	n := 0
	for _, op := range pending {
		if err := wh.Demand(op.Amount); err != nil {
			pending[n] = op
			n++
		} else {
			fmt.Printf("✅ Отложенный заказ выполнен: -%d, Остаток: %d\n",
				op.Amount, wh.balance)
			op.Responce <- true
		}
	}
	return pending[:n]
}
