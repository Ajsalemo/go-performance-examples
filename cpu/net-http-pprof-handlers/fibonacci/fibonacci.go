package fibonacci

import (
	"encoding/json"
	"log"
	"net/http"
)

func Fib(n int) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fib(n - 1) + Fib(n - 2)
	}
}

func Fibonacci(w http.ResponseWriter, r *http.Request) {
	log.Println("Running a Fibonacci function..")
	log.Println(Fib(30))
	m := map[string]string{"msg": "Logging to console.."}
	e := json.NewEncoder(w)
	e.Encode(m)
}