package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Go11Group/at_lesson/lesson34/model"
	"github.com/Go11Group/at_lesson/lesson34/storage/postgres"
	"github.com/gorilla/mux"
)

type Handler struct {
	Talaba *postgres.StudentRepo
}

func NewHandler(repo *postgres.StudentRepo) *Handler {
	return &Handler{Talaba: repo}
}

func (h *Handler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/student/{id}", h.StudentGetById).Methods("GET")
	r.HandleFunc("/student/{id}", h.StudentDeleteById).Methods("DELETE")
	r.HandleFunc("/student/{id}", h.StudentUpdateById).Methods("PUT")
	r.HandleFunc("/student", h.StudentCreate).Methods("POST")
}

func (h *Handler) StudentGetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	student, err := h.Talaba.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func (h *Handler) StudentDeleteById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.Talaba.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("Success to delete student"))
}

func (h *Handler) StudentUpdateById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid student ID", http.StatusBadRequest)
		return
	}

	var student model.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student.ID = id

	err = h.Talaba.Update(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("Success to update"))
}

func (h *Handler) StudentCreate(w http.ResponseWriter, r *http.Request) {
	var student model.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.Talaba.Create(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}
