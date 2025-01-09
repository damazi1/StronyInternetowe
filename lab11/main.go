package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func stronaFunc(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "lab11/index.html")
}

type Student struct {
	Imie     string
	Nazwisko string
	Index    int
	Mail     string `json:"-"`
}

var studenci []Student = []Student{
	{"Jan", "Kowalski", 12345, "test@test"},
	{"Marek", "Nowak", 30000, "to@tamto"},
	{"Anna", "Zdyb", 23232, "anna@zdyb"},
}

func parseFunc(w http.ResponseWriter, r *http.Request) {
	// zwrócenie strony o dynamicznej zawartości
	tmpl, _ := template.ParseFiles("lab11/index.html")
	err := tmpl.Execute(w, studenci)
	if err != nil {
		fmt.Println(err)
		return
	}
}
func ErrorH(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("lab11/err.html")
	err := tmpl.Execute(w, r.PathValue("index"))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func saveMar(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("index")
	indeks, _ := strconv.Atoi(id)
	for i := 0; i < len(studenci); i++ {
		if studenci[i].Index == indeks {
			data, _ := json.Marshal(studenci[i])
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		}
	}
	//w.WriteHeader(http.StatusNotFound)
	http.Redirect(w, r, "/err/{id}", http.StatusSeeOther)
}
func main() {
	http.HandleFunc("/strona/", stronaFunc)
	http.HandleFunc("/index/", parseFunc)
	http.HandleFunc("/save/{index}", saveMar)
	http.HandleFunc("/err/{index}", ErrorH)
	http.ListenAndServe("localhost:8080", nil)
}
