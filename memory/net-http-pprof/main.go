package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/azureossd/go-performance-examples/net-http-pprof/cityarray"
	"github.com/azureossd/go-performance-examples/net-http-pprof/index"
)

func main() {
	http.HandleFunc("/", index.Index)
	http.HandleFunc("/api/array", cityarray.CityArray)

	log.Println("Server listening on port 8080.")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
