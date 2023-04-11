package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Строка подключения к базе данных MySQL в контейнере Docker
	db, err := sql.Open("mysql", "root:417149@tcp(localhost:3306)/www")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Проверяем соединение с базой данных
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Успешное подключение к базе данных MySQL в контейнере Docker!")

	// Insert, err := db.Query("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('title2', 'anons2', 'full_text2')")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("InsertIntoBD")
	// defer Insert.Close()

	// Выполнение SQL-запроса
	rows, err := db.Query("SELECT * FROM articles")
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
