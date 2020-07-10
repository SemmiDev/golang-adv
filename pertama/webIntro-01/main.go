package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "Sammi Aldhi Yanto",
			"Message": "Have a nice day",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		t.Execute(w, data)
	})

	fmt.Println("starting web server at :8080")
	http.ListenAndServe(":8080", nil)
}
