package main

import (
	"net/http"
	"text/template"
)

type page struct {
	Title string //имена переменных должны совпадать с тем, что мы написали выше!
	Msg   string //и переменные обязательно должны быть публичными!
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html")

	t, _ := template.ParseFiles("template/index.html")
	t.Execute(w, &page{Title: "Just page", Msg: "Hello World"})
	
}

func main() {
	http.HandleFunc("/", index)

	http.ListenAndServe(":80", nil)

}
