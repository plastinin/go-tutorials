package main

import (
	"fmt"
	"net/http"
)

func emptyHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello empty!"))
	if err != nil {
		fmt.Println("Ошибка", err.Error())
	} else {
		fmt.Println("ok.")
	}
}

func paylHandler(w http.ResponseWriter, r *http.Request) {
	str := "Pay success."
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Ошибка", err.Error())
	} else {
		fmt.Println("ok.")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Pay canceled."
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Ошибка", err.Error())
	} else {
		fmt.Println("ok.")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	// this is the new gourutine
	str := "Hello world"
	b := []byte(str)

	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Ошибка", err.Error())
	} else {
		fmt.Println("ok.")
	}
}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", paylHandler)
	http.HandleFunc("/cancel", cancelHandler)
	http.HandleFunc("/", emptyHandler)

	fmt.Println("Try to start http...")

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Ошибка:", err.Error())
	}

	fmt.Println("Http server stoped")
}