package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var Db *sql.DB
var err error
var Insert *sql.Rows
var Res *sql.Rows

// подключение к базе данных
func Connect() {
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/www")
	if err != nil {
		panic(err)
	}

	fmt.Println("Connect")
}

// добавление пользователей
func InsertIntoBD() {
	Insert, err = Db.Query("INSERT INTO `user` (`name`, `age`) VALUES ('Max', 44)")
	if err != nil {
		panic(err)
	}
	fmt.Println("InsertIntoBD")
}

// читаем и выводим данные с базы данных
func Read() {
	Res, err = Db.Query("SELECT * FROM `user`")
	if err != nil {
		panic(err)
	}

	for Res.Next() {
		var user User
		err = Res.Scan(&user.Id, &user.Name, &user.Age)
		fmt.Println(user.Id, user.Name, user.Age)
	}
}

func main() {

	Connect()
	defer Db.Close()

	// InsertIntoBD()
	// defer Insert.Close()

	Read()
	defer Res.Close()
}
