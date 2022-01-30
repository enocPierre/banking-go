package main

import (
	"fmt"
	"net/http"
)

func main () {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello word!!")
	})

	http.ListenAndServe("localhost:8080", nil)
}