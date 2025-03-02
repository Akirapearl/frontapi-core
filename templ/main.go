package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := hello("Me")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on port :8000")
	http.ListenAndServe(":8000", nil)
}
