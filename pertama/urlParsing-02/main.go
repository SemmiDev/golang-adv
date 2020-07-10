package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "http://example.com/introduce?name=sam&age=18"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)
	fmt.Printf("host: %s\n", u.Host)
	fmt.Printf("path: %s\n", u.Path)
	fmt.Println("hostname: ", u.Hostname)

	var name = u.Query()["name"][0] // sam
	var age = u.Query()["age"][0]   // 18
	fmt.Printf("name: %s, age: %s\n", name, age)
}
