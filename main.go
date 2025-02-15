package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Disc struct {
	Title string
	Group string
}

func main() {
	fmt.Println("Hi")

	handler1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		dsks := map[string][]Disc{
			"Discs": {
				{Title: "From Zero", Group: "Linkin Park"},
				{Title: "Scoring the End of the World", Group: "Motionless in White"},
			},
		}
		tmpl.Execute(w, dsks)
	}
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
