package server

import (
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Start server
func Start() {

	logrus.Printf("logrus start")

	http.HandleFunc("/", Main)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func Main(w http.ResponseWriter, r *http.Request) {

	logrus.Printf("logrus func main")

	tmpl, err := template.ParseFiles("template/index.html", "template/header.html", "template/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
