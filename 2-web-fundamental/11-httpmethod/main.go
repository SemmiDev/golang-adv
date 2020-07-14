package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Write([]byte("post method"))
		case "GET":
			w.Write([]byte("get method"))
		case "PUT":
			w.Write([]byte("put method"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("started")
	http.ListenAndServe(":9000", nil)
}
