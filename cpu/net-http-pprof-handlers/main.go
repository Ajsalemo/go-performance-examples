package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/azureossd/go-performance-examples/net-http-pprof-handlers/fibonacci"
	"github.com/azureossd/go-performance-examples/net-http-pprof-handlers/index"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", index.Index)
	r.HandleFunc("/api/fib", fibonacci.Fibonacci)
	// Since were using a custom serveMux, we need to register these handles manually for net/http/pprof
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)

	log.Println("Server listening on port 8080.")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}

}
