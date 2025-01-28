package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Web services are easy with Go!")
	})

	http.ListenAndServe("9090", nil)
}
