package main

import (
	"html/template"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/main.html")
	t.Execute(w, nil)
}

func resultsPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/results.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/results", resultsPage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("LAS err: ", err)
	}
}
