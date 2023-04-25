package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Words struct {
	Rus  string
	Norg string
}

var Word1 Words

func main() {
	Word1.Rus = "машина"
	Word1.Norg = "bil"

	log.Println("started http.ListenAndServe localhost:8080")
	http.HandleFunc("/index", index)
	http.HandleFunc("/povei", povei)
	http.HandleFunc("/words", words)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func povei(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/povei.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, Word1.Rus)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func words(w http.ResponseWriter, r *http.Request) {
	WordTmp := r.FormValue("word")
	fmt.Println(WordTmp)

	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
