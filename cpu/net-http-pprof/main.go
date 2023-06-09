package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/azureossd/go-performance-examples/cpu/net-http-pprof/fibonacci"
	"github.com/azureossd/go-performance-examples/cpu/net-http-pprof/index"
)

func main() {
	http.HandleFunc("/", index.Index)
	http.HandleFunc("/api/fib", fibonacci.Fibonacci)

	log.Println("Server listening on port 8080.")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
