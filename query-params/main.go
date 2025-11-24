package main

import (
	"fmt"
	"net/http"
)

func getHandle(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.URL.Query() {
		fmt.Printf("%s: %s\n", k, v)
	}
}

func main() {

	fmt.Println("Server started...")
	http.HandleFunc("/get", getHandle)
	http.ListenAndServe(":9091", nil)

}
