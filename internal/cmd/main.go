package main

import (
	"fmt"
	"net/http"

	"github.com/gh0st3e/TodoForTechArt/internal/repository"
)

func handleRequest() {

	http.HandleFunc("/", repository.Index)
	http.HandleFunc("/save_todo", repository.AddTodo)
	http.HandleFunc("/del_todo", repository.DelTodo)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is Work")

}

func main() {
	handleRequest()
}
