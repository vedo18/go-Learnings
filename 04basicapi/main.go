package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{"text": "Hello from Go server!!!"}

	json.NewEncoder(w).Encode(response)
}

func main() {
	fmt.Println("Welcome to go server!")
	http.HandleFunc("/get", getHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
	// fmt.Println("Server listening on", http.ListenAndServe(":8080", nil)) //another way to listen
}
