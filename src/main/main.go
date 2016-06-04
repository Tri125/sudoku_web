package main

import (
	"github.com/gorilla/mux"
	"github.com/tri125/sudoku"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var t *template.Template = ParseTemplates()

func main() {
	var port int = 4040
	portString := strconv.Itoa(port)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/", SolveHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", r)

	log.Print("Listening on port " + portString)

	err := http.ListenAndServe(":"+portString, nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}

}

func ParseTemplates() *template.Template {
	t := template.New("home.html").Funcs(template.FuncMap{
		"loop": func(n int) []struct{} {
			return make([]struct{}, n)
		},
	}).Funcs(template.FuncMap{
		"each": func(interval int, n int) bool {
			return n%interval == 0
		},
	})

	t, err := t.ParseFiles("templates/home.html")

	if err != nil {
		log.Print("template/home error:", err)
	}

	return t
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {
	var gridValue [81]string
	t.ExecuteTemplate(w, "home.html", gridValue)
}

func SolveHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	var gridPost []string

	if err != nil {
		log.Print("SolveHandler ParseForm error:", err)
	}

	for key, values := range req.PostForm {
		if key == "value[]" {
			gridPost = values
		} else {
			continue
		}

	}
	log.Print("AH:", gridPost[0] == "")
	if gridPost != nil {
		var grid sudoku.Grid
		var count int = 0
		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				gridValue, err := strconv.Atoi(gridPost[count])
				if err != nil {
					log.Print("Post grid atoi error:", err)
				}
				grid[x][y] = gridValue
				count++
			}
		}
		grid, err := sudoku.SolveGrid(grid)
		if err != nil {
			log.Print("Sudoku Grid Solver error:", err)
		}
		//jsonResponse, err := json.Marshal(grid)

		if err != nil {
			log.Print("Error:", err)
		}
		var flatGrid [81]int
		count = 0
		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[x]); y++ {
				flatGrid[count] = grid[x][y]
				count++
			}
		}
		//w.Write(jsonResponse)
		t.ExecuteTemplate(w, "home.html", flatGrid)
	}

}
