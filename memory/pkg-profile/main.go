package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/azureossd/go-performance-examples/memory/profile/cityarray"
	"github.com/azureossd/go-performance-examples/memory/profile/index"
	"github.com/google/uuid"
	"github.com/pkg/profile"
)

func main() {
	id := uuid.New().String()
	cdir, _ := os.Getwd()
	path := filepath.Join(cdir, id)
	// This creates a file named mem.pprof in a randomly named directory (by UID) relative to the current directory
	defer profile.Start(profile.MemProfile, profile.ProfilePath(path)).Stop()
	http.HandleFunc("/", index.Index)
	http.HandleFunc("/api/array", cityarray.CityArray)

	log.Println("Server listening on port 8080.")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
