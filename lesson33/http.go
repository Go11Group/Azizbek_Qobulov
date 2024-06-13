package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api", apiHandler)
	http.HandleFunc("/api/resource", resourceHandler)

	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getHandler(w, r)
	case "POST":
		postHandler(w, r)
	case "PUT":
		putHandler(w, r)
	case "DELETE":
		deleteHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	w.Write([]byte(fmt.Sprintf("GET request received with param: %s\n", param)))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	w.Write([]byte(fmt.Sprintf("POST request received with body: %s\n", string(body))))
}


func putHandler(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	w.Write([]byte(fmt.Sprintf("PUT request received with body: %s\n", string(body))))
}


func deleteHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("param")
	w.Write([]byte(fmt.Sprintf("DELETE request received with param: %s\n", param)))
}


func resourceHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		w.Write([]byte(fmt.Sprintf("Resource GET request received for id: %s\n", id)))
	case "POST":
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		w.Write([]byte(fmt.Sprintf("Resource POST request received for id: %s with body: %s\n", id, string(body))))
	case "PUT":
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		w.Write([]byte(fmt.Sprintf("Resource PUT request received for id: %s with body: %s\n", id, string(body))))
	case "DELETE":
		w.Write([]byte(fmt.Sprintf("Resource DELETE request received for id: %s\n", id)))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
