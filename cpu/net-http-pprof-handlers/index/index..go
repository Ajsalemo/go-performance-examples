package index

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	m := map[string]string{"msg": "go-performance-examples"}
	e := json.NewEncoder(w)
	e.Encode(m)
}

