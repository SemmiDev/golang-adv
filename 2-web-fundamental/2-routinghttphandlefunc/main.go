package main

import (
	"fmt"
	"net/http"
)

// method
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index"))
}

func main() {
	// closure
	samhandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello sam"))
	}

	http.HandleFunc("/", samhandler)
	http.HandleFunc("/index", indexHandler)

	// anonymous
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello dev"))
	})

	// running server
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
