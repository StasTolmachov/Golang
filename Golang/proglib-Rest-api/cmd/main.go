package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Golang/proglib-Rest-api/internal/model"
	"github.com/Golang/proglib-Rest-api/internal/store"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()
	sugar.Infow("Started: http://127.0.0.1:8000/books")

	r := mux.NewRouter()

	model.Books = append(model.Books, model.Book{ID: "1", Title: "Война и Мир", Author: &model.Author{Firstname: "Лев", Lastname: "Толстой"}})
	model.Books = append(model.Books, model.Book{ID: "2", Title: "Преступление и наказание", Author: &model.Author{Firstname: "Фёдор", Lastname: "Достоевский"}})

	r.HandleFunc("/books", store.GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", store.GetBook).Methods("GET")
	r.HandleFunc("/books", store.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", store.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", store.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
