package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ActionIndexJsonEncoder(response http.ResponseWriter, request *http.Request) {
	data := []struct {
		Name string
		Age  int
	}{
		{"Sam", 19},
		{"Rauf", 19},
		{"Adit", 19},
		{"Dandi", 19},
		{"Gusnur", 19},
		{"Ayatullah", 50},
	}
	err := json.NewEncoder(response).Encode(data)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("content-type", "application/json")
}

func ActionIndexMarshal(response http.ResponseWriter, request *http.Request) {
	datas := []struct {
		Name string
		Age  int
	}{
		{"Sam", 19},
		{"Rauf", 19},
		{"Adit", 19},
		{"Dandi", 19},
		{"Gusnur", 19},
		{"Ayatullah", 50},
	}
	err := json.NewEncoder(response).Encode(datas)

	jsonInBytes, err := json.Marshal(datas)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(jsonInBytes)
}

func main() {
	http.HandleFunc("/jsonencoder", ActionIndexJsonEncoder)
	http.HandleFunc("/jsonmarshal", ActionIndexMarshal)
	fmt.Println("starting app in :9000")
	http.ListenAndServe(":9000", nil)
}
