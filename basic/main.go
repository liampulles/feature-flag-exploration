package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting")
	err := http.ListenAndServe(":9080", http.HandlerFunc(handler))
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
	}
	fmt.Println("closing")
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hi there!!!\n"))
}
