<<<<<<< HEAD
package handler

import (
	"database/sql"

	"github.com/Go11Group/at_lesson/lesson34/model"
)

type StudentRepo struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{db: db}
}

func (r *StudentRepo) GetById(id string) (*model.Student, error) {
	query := "SELECT id, name, phone, email FROM students WHERE id = $1"
	row := r.db.QueryRow(query, id)
	var student model.Student
	if err := row.Scan(&student.ID, &student.Name, &student.Phone, &student.Email); err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepo) Delete(id string) error {
	query := "DELETE FROM students WHERE id = $1"
	_, err := r.db.Exec(query, id)
	return err
}

func (r *StudentRepo) Update(student *model.Student) error {
	query := "UPDATE students SET name = $1, phone = $2, email = $3 WHERE id = $4"
	_, err := r.db.Exec(query, student.Name, student.Phone, student.Email, student.ID)
	return err
}

func (r *StudentRepo) Create(student *model.Student) error {
	query := "INSERT INTO students (name, phone, email) VALUES ($1, $2, $3) RETURNING id"
	return r.db.QueryRow(query, student.Name, student.Phone, student.Email).Scan(&student.ID)
}
=======
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
>>>>>>> origin/main
