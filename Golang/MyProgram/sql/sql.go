package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Articles struct {
	Id        int
	Title     string
	Full_text string
}

func main() {
	fmt.Println("open start")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/golang")
	if err != nil {
		panic(err)
	}
	fmt.Println("open finish")
	// defer db.Close()

	res, err := db.Query("SELECT * FROM `articles`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var articles Articles
		err = res.Scan(&articles.Id, &articles.Title, &articles.Full_text)
		fmt.Println(articles)
	}

}
