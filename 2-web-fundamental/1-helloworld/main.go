package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	var message = "hello sammidev"
	w.Write([]byte(message))
}

func timenow(w http.ResponseWriter, r *http.Request) {
	var now = time.Now()
	w.Write([]byte(strconv.Itoa(now.Year())))
}

func main() {
	http.HandleFunc("/", greeting)
	http.HandleFu
	http.HandleFunc("/now", timenow)

	var address = ":9000"
	fmt.Printf("server started at %s\n", address)

	// cara pertama bikin web server
	// err := http.ListenAndServe(address, nil)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// cara kedua
	server := new(http.Server)
	server.Addr = address
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
