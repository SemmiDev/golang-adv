package main

import (
	"fmt"
	"net/http"
)

func main() {
	// routing static
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("started at :9000")
	http.ListenAndServe(":9000", nil)
}
