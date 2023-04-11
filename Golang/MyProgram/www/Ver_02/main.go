package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
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
	// title := r.FormValue("title") возвращает пустую строку. как испраить?
	title := r.FormValue("title")
	if title == "" {
		title = "Default Title"
	}
	anons := r.FormValue("anons")
	if anons == "" {
		anons = "Default anons"
	}

	full_text := r.FormValue("full_text")
	if full_text == "" {
		full_text = "Default full_text"
	}

	Db, err := sql.Open("mysql", "root:417149@tcp(localhost:3306)/www")
	if err != nil {
		panic(err)
	}
	defer Db.Close()

	Insert, err := Db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s')", title, anons, full_text))
	// Insert, err := Db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `full_text`) VALUES ('%s', `%s`)", title, full_text))
	if err != nil {
		panic(err)
	}
	defer Insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)

	rows, err := Db.Query("SELECT * FROM articles")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Обработка результатов запроса
	for rows.Next() {
		var id int
		var title string
		var anons string
		var full_text string
		err = rows.Scan(&id, &title, &anons, &full_text)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("id: %v, title: %s, anons: %s, full_text: %s\n", id, title, anons, full_text)
	}

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
