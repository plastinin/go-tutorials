package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

var money = 1000
var bank = 0
var mtx = sync.Mutex{}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Read IO error:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("Failed to convert", err)
		return
	}
	mtx.Lock()
	if money - paymentAmount >= 0 {
		money -= paymentAmount
		fmt.Println("Pay success:", paymentAmount, "USD. Balance:", money)
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Read IO error:", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("Failed to convert", err)
		return
	}

	mtx.Lock()
	if money >= saveAmount {
		money -= saveAmount
		bank +=saveAmount
		fmt.Println("Bank success:", saveAmount, "USD. Bank Balance:", bank, "Money:", money)
	} else {
		fmt.Println("Not enought money")
	}
	mtx.Unlock()
}

func main() {

	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler)

	fmt.Println("Server start")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

}