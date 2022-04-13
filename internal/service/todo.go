package service

import (
	"github.com/gh0st3e/TodoForTechArt/internal/repository"
	"net/http"
)

type todo struct {
	tRepo repository.Todo
}

func (to todo) Index(w http.ResponseWriter, r *http.Request) {
	to.tRepo.Index(w, r)
}
