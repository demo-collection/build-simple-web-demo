package main

import (
	"../core"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", core.SayHelloName)
	http.HandleFunc("/login", core.Login)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal(err)
	}
}
