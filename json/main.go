package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Payment struct {
	Description string `json:"description"` // описание покупки
	USD         int    `json:"USD"`         // сумма покупк
	FullName    string `json:"fullName"`    // ФИО
	Address     string `json:"address"`     // Адрес
	Time        time.Time
}

type HttpResponse struct {
	Money          int
	PaymentHistory []Payment
}

func (p Payment) Println() {
	fmt.Println("description:", p.Description)
	fmt.Println("USD:", p.USD)
	fmt.Println("fullName:", p.FullName)
	fmt.Println("address:", p.Address)
}

var money = 1000
var paymentHistory = make([]Payment, 0)
var mtx = sync.Mutex{}

func payHandler(w http.ResponseWriter, r *http.Request) {

	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		fmt.Println("err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	payment.Time = time.Now()

	mtx.Lock()
	payment.Println()

	if money-payment.USD >= 0 {
		money -= payment.USD
	}
	paymentHistory = append(paymentHistory, payment)

	HttpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	b, err := json.MarshalIndent(HttpResponse, "", "	")
	if err != nil {
		fmt.Println("Error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(b); err != nil {
		fmt.Println("err:", err)
		return
	}
	mtx.Unlock()

}

func main() {
	http.HandleFunc("/pay", payHandler)
	fmt.Println("Starting http-service...")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Error connect:", err.Error())
	}
}
