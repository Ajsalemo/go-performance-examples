package main

import (
	"log"
	_ "net/http/pprof"

	"github.com/azureossd/go-deployment-samples/index"
	"github.com/azureossd/go-deployment-samples/fibonacci"
)

func main() {
	http.HandleFunc("/", index.Cars)
	http.HandleFunc("/api/fib", fibonacci.Fibonacci)

	log.Println("Server listening on port 8080.")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
