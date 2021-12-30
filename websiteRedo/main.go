package main

import (
	"html/template"
	"net/http"
)

type GData struct {
	Title string
	Sen1  string
	Sen2  string
	Sen3  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.Handle("/assets/",
		http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8081", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Title: "Welcome to my sample Golang Website!",
	}
	tpl.ExecuteTemplate(w, "index.gohtml", gd)
}

func about(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Sen1: "Hello my name is Garcia Reynaldo Milord welcome to my Website. Am college graduate at Florida International University ",
		Sen2: "BS in Computer Science, 2019-2021. Interested in software development, rest API, and AI/ML.",
		Sen3: "Hobbies include reading comic, fitness, and eating delicious food. ",
	}
	tpl.ExecuteTemplate(w, "about.gohtml", gd)
}

func contact(w http.ResponseWriter, r *http.Request) {
	gd := GData{
		Title: "Contact",
	}
	tpl.ExecuteTemplate(w, "contact.gohtml", gd)
}
