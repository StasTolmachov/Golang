package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

type articles struct {
	Id                      int
	Title, Anons, Full_text string
}

var posts = []articles{}
var showPost = articles{}
var NokUah = "3.445"

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Db, err := sql.Open("mysql", "root:417149@tcp(localhost:3306)/www")
	if err != nil {
		panic(err)
	}
	defer Db.Close()

	rows, err := Db.Query("SELECT * FROM articles")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	posts = []articles{}

	// Обработка результатов запроса
	for rows.Next() {
		var news articles
		err = rows.Scan(&news.Id, &news.Title, &news.Anons, &news.Full_text)
		if err != nil {
			panic(err.Error())
		}

		posts = append(posts, news)
	}

	err = t.ExecuteTemplate(w, "index", posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/create.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "create", nil)

}
func createTest(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/createTest.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "createTest", nil)
	// t.Execute(w, NokUah)

}

func safe_article(w http.ResponseWriter, r *http.Request) {
	fmt.Println("safe_article")

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

	if err != nil {
		panic(err)
	}
	defer Insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func show_post(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/show.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Db, err := sql.Open("mysql", "root:417149@tcp(localhost:3306)/www")
	if err != nil {
		panic(err)
	}
	defer Db.Close()

	rows, err := Db.Query(fmt.Sprintf("SELECT * FROM articles WHERE id = %s", vars["id"]))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	showPost = articles{}

	// Обработка результатов запроса
	for rows.Next() {
		var post articles
		err = rows.Scan(&post.Id, &post.Title, &post.Anons, &post.Full_text)
		if err != nil {
			panic(err.Error())
		}

		showPost = post
	}

	err = t.ExecuteTemplate(w, "show", showPost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func handleFunc() {
	fmt.Println("run")

	rtr := mux.NewRouter()

	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/create/", create).Methods("GET")
	rtr.HandleFunc("/createTest/", createTest).Methods("GET")
	rtr.HandleFunc("/safe_article/", safe_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_post).Methods("GET")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.ListenAndServe(":8080", nil)

}



func main() {

	handleFunc()

	fmt.Println(NokUah)
}
