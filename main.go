package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	if err := openDB(); err != nil {
		log.Printf("ERROR connecting to database %v", err)											
	}
	defer closeDB()

	tmpl := template.Must(template.ParseFiles("templates/index.html"))      						
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))      

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data, err := getDashboardData()
		if err != nil {
			http.Error(w, "Error fetching data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

    fmt.Println("Server is listening on port 8080...")												
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}