package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	client := &http.Client{}

	// Yangi so'rov yaratamiz va uni 8080-portga yuboramiz
	req, err := http.NewRequest(r.Method, "http://localhost:8080"+r.RequestURI, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header = r.Header

	res, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	// Natijani o'qiymiz va chop etamiz
	body, err := io.ReadAll(res.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Response from server 8080:", string(body))

	// Javobni klientga yuboramiz
	for k, v := range res.Header {
		w.Header()[k] = v
	}
	w.WriteHeader(res.StatusCode)
	w.Write(body)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 7070")
	if err := http.ListenAndServe(":7070", nil); err != nil {
		fmt.Println(err)
	}
}
