package main

import (
	"log"
	"net/http"

	"./handlers"
)

func handleRequest() {
	http.HandleFunc("/organisation/accounts", handlers.AccountRouter)
	http.HandleFunc("/organisation/accounts/", handlers.AccountRouter)
	http.HandleFunc("/", handlers.RootHandler)
}

func main() {
	handleRequest()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
