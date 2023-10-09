package main

import (
	"app-htmx/routing"
	"net/http"
)

func main() {
	http.HandleFunc("/", routing.FrontController)
	http.ListenAndServe(":5000", nil)
}
