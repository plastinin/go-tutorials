package common

type Operation struct {
	Amount   int
	Responce chan bool // статус исполнения заказа
}
