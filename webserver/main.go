package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Message struct {
	Content string `json:"content"`
}

var (
	messages  = make(map[int]Message)
	idCounter = 1
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing 'id' parameter")
		return
	}

	for key, msg := range messages {
		if fmt.Sprintf("%d", key) == id {
			json.NewEncoder(w).Encode(msg)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Message not found")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}
	defer r.Body.Close()

	var msg Message
	if err := json.Unmarshal(body, &msg); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid JSON")
		return
	}

	messages[idCounter] = msg
	idCounter++
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing 'id' parameter")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}
	defer r.Body.Close()

	var msg Message
	if err := json.Unmarshal(body, &msg); err != nil {
		fmt.Fprintf(w, "Invalid JSON")
	}

	for key := range messages {
		if fmt.Sprintf("%d", key) == id {
			messages[key] = msg
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(msg)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Message not found")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Missing 'id' parameter")
		return
	}

	for key := range messages {
		if fmt.Sprintf("%d", key) == id {
			delete(messages, key)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Message deleted")
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Message not found")
}

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/put", putHandler)
	http.HandleFunc("/delete", deleteHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
