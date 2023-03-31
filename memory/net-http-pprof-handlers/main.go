package main

import (
	"log"
	"net/http"
	"net/http/pprof"

	"github.com/azureossd/go-performance-examples/memory/net-http-pprof-handlers/cityarray"
	"github.com/azureossd/go-performance-examples/memory/net-http-pprof-handlers/index"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", index.Index)
	r.HandleFunc("/api/fib", cityarray.CityArray)
	// Since were using a custom serveMux, we need to register these handles manually for net/http/pprof
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)

	log.Println("Server listening on port 8080.")
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}

}
