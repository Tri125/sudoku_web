package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var port int = 4040
	portString := strconv.Itoa(port)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", r)

	log.Print("Listening on port " + portString)

	err := http.ListenAndServe(":"+portString, nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}

}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	t := template.New("home.html").Funcs(template.FuncMap{
		"loop": func(n int) []struct{} {
			return make([]struct{}, n)
		},
	})
	t, err := t.ParseFiles("templates/home.html")

	if err != nil {
		log.Print("template/home error:", err)
	}
	t.ExecuteTemplate(w, "home.html", nil)
}
