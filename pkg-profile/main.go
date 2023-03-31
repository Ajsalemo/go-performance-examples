package main

import (
	"log"
	"net/http"

	"github.com/pkg/profile"

	"github.com/azureossd/go-performance-examples/profile/fibonacci"
	"github.com/azureossd/go-performance-examples/profile/index"
)

func main() {
	// This creates a file named cpu.pprof in the current directory
	defer profile.Start(profile.ProfilePath(".")).Stop()
	http.HandleFunc("/", index.Index)
	http.HandleFunc("/api/fib", fibonacci.Fibonacci)

	log.Println("Server listening on port 8080.")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}
