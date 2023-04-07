package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"net/http"
	"text/template"
)



func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)
}



func safe_article(w http.ResponseWriter, r *http.Request) {
	fmt.Println("safe_article")
	title := r.FormValue("title")
	full_text := r.FormValue("full_text")

	Db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
	if err != nil {
		panic(err)
	}
	defer Db.Close()

	Insert, err := Db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `full_text`) VALUES ('%s', `%s`)", title, full_text))
	if err != nil {
		panic(err)
	}
	defer Insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handleFunc() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create/", create)
	http.HandleFunc("/safe_article/", safe_article)
	http.ListenAndServe(":8080", nil)

}

func main() {
	fmt.Println("run")
	handleFunc()

}
