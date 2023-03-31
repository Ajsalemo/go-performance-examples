package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/azureossd/go-performance-examples/profile/fibonacci"
	"github.com/azureossd/go-performance-examples/profile/index"
	"github.com/google/uuid"
	"github.com/pkg/profile"
)

func main() {
	id := uuid.New().String()
	cdir, _ := os.Getwd()
	path := filepath.Join(cdir, id)
	// This creates a file named cpu.pprof in the current directory
	defer profile.Start(profile.ProfilePath(path)).Stop()
	http.HandleFunc("/", index.Index)
	http.HandleFunc("/api/fib", fibonacci.Fibonacci)

	log.Println("Server listening on port 8080.")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
