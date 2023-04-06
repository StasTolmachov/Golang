package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
}

func (u User) getAllInfo() string {
	return fmt.Sprintf("User is: %s%d%d", u.Name, u.Age, u.Money)
}

func (u *User) setNewName(newName string) {
	u.Name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{Name: "Bob", Age: 25, Money: -50, Avg_grades: 4.2, Happiness: 0.8}
	// bob.setNewName("Alex")
	// fmt.Fprintf(w, bob.getAllInfo())

	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contacts_page")
}

func features_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "features_page")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.HandleFunc("/features/", features_page)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
