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
	
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	} else {
		fmt.Println("Method: ", r.Method)
	}

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "Read IO error:" + err.Error()
		fmt.Println(msg)
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("failed to write HTTP-response:", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		msg := "Failed to convert:" + err.Error()
		fmt.Println(msg)
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("failed to write HTTP-response:", err)
		}
		return
	}
	mtx.Lock()
	if money - paymentAmount >= 0 {
		money -= paymentAmount
		msg := "Pay success: " + strconv.Itoa(paymentAmount) + " USD. Balance: " + strconv.Itoa(money)
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("failed to write HTTP-response:", err)
		}
	}
	mtx.Unlock()
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "Read IO error:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("failed to write HTTP-response:", err)
		}
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