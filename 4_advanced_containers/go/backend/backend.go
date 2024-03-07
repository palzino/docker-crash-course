package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RequestData struct {
	Name string `json:"name"`
}

func main() {
	fmt.Println("Server starting...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Backend!")
	})

	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/get", getHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var requestData RequestData
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		http.Error(w, "Error parsing JSON request body", http.StatusBadRequest)
		return
	}

	// Ensure the /data directory exists
	err = os.MkdirAll("/data", os.ModePerm)
	if err != nil {
		http.Error(w, "Error creating data directory", http.StatusInternalServerError)
		return
	}

	// Marshal requestData to JSON for storage
	dataToSave, err := json.Marshal(requestData)
	if err != nil {
		http.Error(w, "Error marshaling request data", http.StatusInternalServerError)
		return
	}

	// Write the JSON data to /data/requests.json
	err = os.WriteFile("/data/requests.json", dataToSave, 0644)
	if err != nil {
		http.Error(w, "Error writing to file", http.StatusInternalServerError)
		return
	}

	// Assuming you want to send a JSON response back
	response := map[string]string{"message": "Data received successfully", "name": requestData.Name}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error generating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
func getHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	data, err := os.ReadFile("/data/requests.json")
	if err != nil {
		http.Error(w, "Error reading from file", http.StatusInternalServerError)
		return
	}

	if len(data) == 0 {
		http.Error(w, "File is empty", http.StatusNotFound)
		return
	}

	// Wrap the raw data in a JSON object, assuming the raw data is a JSON string
	// If the data structure is known and you want to unmarshal, modify, then remarshal,
	// you would unmarshal `data` into a struct, then marshal that struct into JSON.

	// For simplicity, assuming we're just echoing the raw JSON content
	w.Header().Set("Content-Type", "application/json")
	w.Write(data) // `data` is assumed to be a valid JSON format already
}
