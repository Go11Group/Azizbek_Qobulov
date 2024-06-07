package handler

import (
	"net/http"

	"github.com/Go11Group/at_lesson/lesson34/storage/postgres"
)

type Handler struct {
	Talaba *postgres.StudentRepo
}

func NewHandler(handler Handler) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/student/", handler.student)

	return &http.Server{Handler: mux}
}
