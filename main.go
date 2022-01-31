package main

import (
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTask []task

var task = allTask{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Api")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.hadleFunc("/", indexRoute)
}
