package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// return mux/router
	r := echo.New()

	// context = include responsewriter and req
	// plaintext
	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from root"
		return ctx.String(http.StatusOK, data)
	})

	// html
	r.GET("/html", func(ctx echo.Context) error {
		data := "Hello from /html"
		return ctx.HTML(http.StatusOK, data)
	})

	// redirect
	r.GET("/home", func(ctx echo.Context) error {
		return ctx.Redirect(http.StatusTemporaryRedirect, "/")
	})

	// json
	r.GET("/json", func(ctx echo.Context) error {
		data := M{"Message": "Hello sam", "Counter": 2}
		return ctx.JSON(http.StatusOK, data)
	})

	// PARSING REQUEST
	// query param
	r.GET("/page1", func(ctx echo.Context) error {
		name := ctx.QueryParam("name")
		data := fmt.Sprintf("hello %s", name)
		return ctx.String(http.StatusOK, data)
	}) // testing -> curl -X GET http://localhost:9000/page1?name=sam

	// path and next
	r.GET("/page2/:name/*", func(ctx echo.Context) error {
		name := ctx.Param("name")
		message := ctx.Param("*")

		data := fmt.Sprintf("hello %s i have message for u: %s", name, message)
		return ctx.String(http.StatusOK, data)
	}) // testing -> curl -X GET http://localhost:9000/page2/tim/need/some/sleep

	// parsing form data
	r.POST("/page4", func(ctx echo.Context) error {
		name := ctx.FormValue("name")
		message := ctx.FormValue("message")

		data := fmt.Sprintf(
			"Hello %s, I have message for you: %s",
			name,
			strings.Replace(message, "/", "", 1),
		)

		return ctx.String(http.StatusOK, data)
	}) // testing -> curl -X POST -F name=sam -F message=angry http://localhost:9000/page4

	// echo wrap handler
	var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from action index"))
	}

	// wrap
	var ActionIndex = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("from action index"))
	}

	var ActionHome = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("from action home"))
		},
	)

	var ActionAbout = echo.WrapHandler(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("from action about"))
			},
		),
	)

	// static assets
	// make a folder = assets
	// make a file layout.js
	r.Static("/static", "assets")
	// testing -> http://localhost:9000/static/layout.js

	r.GET("/index", echo.WrapHandler(http.HandlerFunc(ActionIndex)))
	r.GET("/home", echo.WrapHandler(ActionHome))
	r.GET("/about", ActionAbout)

	r.Start(":9000")
}
