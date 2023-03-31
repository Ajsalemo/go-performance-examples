package cityarray

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CityArray(w http.ResponseWriter, r *http.Request) {

	for i := 0; i < 100000; i++ {
		s := []int{}
		m := append(s, i)
		fmt.Println(m)
	}

	m := map[string]string{"msg": "Creating arrays.."}
	e := json.NewEncoder(w)
	e.Encode(m)
}
