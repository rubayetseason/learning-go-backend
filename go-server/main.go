package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, I am Season! I am a web developer. I am learning Golang.")
}

func main() {
	mux := http.NewServeMux() // router

	mux.HandleFunc("/hello", helloHandler) // hello route

	mux.HandleFunc("/about-me", aboutHandler) // about-me route

	fmt.Println("🚀 Golang server is running on: 5000")

	err := http.ListenAndServe(":5000", mux) // if no error then will get nill else will get the error text

	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
