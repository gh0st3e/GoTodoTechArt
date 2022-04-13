package service

import "net/http"

type Todo interface {
	Index(w http.ResponseWriter, r *http.Request)
	DelTodo(w http.ResponseWriter, r *http.Request)
	AddTodo(w http.ResponseWriter, r *http.Request)
	NewID()
}
