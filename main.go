package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TodoItem struct {
	Item string `json:"item"`
}

func main() {
	var todos = make([]string, 0)
	mux := http.NewServeMux()

	// Unified Handler for GET and POST requests
	mux.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// Encode todos as JSON and write to the response
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(todos); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else if r.Method == http.MethodPost {
			// Decode the new to-do item from the request body
			var t TodoItem
			if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Append the new item to the todos slice
			todos = append(todos, t.Item)
			w.WriteHeader(http.StatusCreated)
		} else {
			// Method not allowed
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Start the HTTP server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
