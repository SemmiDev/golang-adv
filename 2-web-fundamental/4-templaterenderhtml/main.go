package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "home.html")
	var tmplt, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{}{
		"title": "Learning golang web",
		"name":  "Sam",
	}

	err = tmplt.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// handling file statis
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	// handling root
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var filepath = path.Join("views", "login.html")
		var tmplt, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title": "Learning golang web",
			"name":  "Sam",
		}

		err = tmplt.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("started at server localhost :9000")
	http.ListenAndServe(":9000", nil)
}
