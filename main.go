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
	// Initial page load
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		// Initial data (empty or nil)
		tmpl.Execute(w, nil)
	})

	// Handler for HTMX read request
	http.HandleFunc("/read/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		// Fetch data (replace with database call)
		data := getDiscsFromDB()
		// Render only the partial template
		tmpl.ExecuteTemplate(w, "discs-partial", data)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Simulate database fetch
func getDiscsFromDB() map[string][]Disc {
	return map[string][]Disc{
		"Discs": {
			{Title: "From Zero", Group: "Linkin Park"},
			{Title: "Scoring the End of the World", Group: "Motionless in White"},
		},
	}
}
