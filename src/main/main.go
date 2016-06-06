package main

import (
	"errors"
	"github.com/gorilla/mux"
	lumber "github.com/jcelliott/lumber"
	"github.com/tri125/sudoku"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

//Lumber logger
var logger lumber.Logger

//Parse our templates
var t *template.Template = ParseTemplates()

func main() {
	//Port that the app will listen for requests
	var port int = 4040
	portString := strconv.Itoa(port)

	//Set lumber
	logger, _ = lumber.NewFileLogger("filename.log", lumber.INFO, lumber.ROTATE, 5000, 9, 100)
	//Create the Gorilla Mux
	r := mux.NewRouter()
	//Handler for Root GET
	r.HandleFunc("/", HomeHandler).Methods("GET")
	//Handler for Root POST
	r.HandleFunc("/", SolveHandler).Methods("POST")
	//To be able to serve public files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	http.Handle("/", r)

	log.Print("Listening on port " + portString)

	err := http.ListenAndServe(":"+portString, nil)
	if err != nil {
		log.Fatal("ListenAndServe error:", err)
	}

}

//ParseTemplates return a pointer to a Template type.
//All templates are initialized in this function.
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
		log.Fatal("template/home error:", err)
	}

	return t
}

//HomeHandler handle GET requests on the root directory
func HomeHandler(w http.ResponseWriter, req *http.Request) {
	//String array of 81 elements representing the values for an empty sudoku grid.
	var gridValue [81]string
	t.ExecuteTemplate(w, "home.html", gridValue)
}

//SolveHandler handle POST requests on the root directory
//Responsible to manipulate POST form data and present to the user a solved sudoku grid.
func SolveHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	var gridPost []string

	if err != nil {
		log.Print("SolveHandler ParseForm error:", err)
	}

	//Retrieve the POST data for our "value[]" input elements.
	for key, values := range req.PostForm {
		if key == "value[]" {
			//Contain a string array
			gridPost = values
		} else {
			continue
		}

	}

	if gridPost != nil {
		solvedGrid, err := SolvePost(gridPost)

		if err != nil {
			log.Print("Error handling Post and solving:", err)
		}
		flatGrid := FlattenGrid(solvedGrid)
		t.ExecuteTemplate(w, "home.html", flatGrid)
	}

}

//SolvePost accept an array of string, convert it to a sudoku.Grid type and solve the grid.
//Returns the solved Grid and error if applicable.
func SolvePost(gridPost []string) (answer sudoku.Grid, err error) {
	var grid sudoku.Grid
	var count int = 0
	ch := make(chan sudoku.Grid, 1)

	//Itterate on the 2d array and assign its values.
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			gridValue, err := strconv.Atoi(gridPost[count])
			if err != nil {
				logger.Warn(err.Error())
				log.Print("Post grid atoi error:", err)
			}
			grid[x][y] = gridValue
			count++
		}
	}

	//Solve the grid
	go _solveGrid(grid, ch)

	select {
	case solvedGrid := <-ch:
		return solvedGrid, nil
	case <-time.After(5 * time.Second):
		return grid, errors.New("TimeOut")
	}

	return grid, nil
}

func _solveGrid(grid sudoku.Grid, ch chan sudoku.Grid) {
	grid, err := sudoku.SolveGrid(grid)

	if err != nil {
		log.Print("Sudoku Grid Solver error:", err)
	}
	ch <- grid
}

//FlattenGrid take a sudoku Grid type (2d int array) and flatten it in one dimension.
//Return an int array.
func FlattenGrid(grid sudoku.Grid) (flatGrid [81]int) {
	var count int = 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			flatGrid[count] = grid[x][y]
			count++
		}
	}
	return flatGrid
}
