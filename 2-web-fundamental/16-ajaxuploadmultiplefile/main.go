package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/upload", handleUpload)
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at :7000")
	http.ListenAndServe(":7000", nil)
}

func handleIndex(response http.ResponseWriter, request *http.Request) {
	tmpl := template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(response, nil); err != nil {
		http.Error(response, err.Error(),
			http.StatusInternalServerError)
	}
}

func handleUpload(response http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		http.Error(response, "Only accept POST request", http.StatusBadRequest)
		return
	}

	basePath, _ := os.Getwd()
	reader, err := request.MultipartReader()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		fileLocation := filepath.Join(basePath, "files", part.FileName())
		dst, err := os.Create(fileLocation)
		if dst != nil {
			defer dst.Close()
		}
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(dst, part); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	response.Write([]byte(`all files uploaded`))
}
