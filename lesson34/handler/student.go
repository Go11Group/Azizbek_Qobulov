package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Go11Group/at_lesson/lesson34/model"
)

func (h *Handler) student(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)
	switch r.Method {
	case "GET":
		h.StudentGetById(w, r)
	case "DELETE":
		h.StudentDeleteById(w, r)
	case "PUT":
		h.StudentUpdateById(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func (h *Handler) StudentGetById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/student/")

	book, err := h.Talaba.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while Decode, err: %s", err.Error())))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(book)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error while Encode, err: %s", err.Error())))
	}
}

func (h *Handler) StudentDeleteById(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/student/")

	err := h.Talaba.Delete(id)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error while Delete, err: %s", err.Error())))
		return
	}
	w.Write([]byte("Succes to delete student"))
}

func (h *Handler) StudentUpdateById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/student/"))
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error id, %s", err)))
	}
	var student model.Student

	student.ID = id
	student.Name = "Azizbek"
	student.Phone = "930693005"
	student.Email = "someone@gmail.com"

	err = h.Talaba.Update(&student)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error on update, %s", err)))
	} else {
		w.Write([]byte("Succes to update"))
	}
}
