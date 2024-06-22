// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type DataStore struct {
	sync.RWMutex
	data map[string]string
}

var store = DataStore{
	data: make(map[string]string),
}

func main() {
	http.HandleFunc("/data", addData)
	http.HandleFunc("/data/", getData)
	http.HandleFunc("/health", healthCheck)
	go backgroundTask()
	log.Println("Starting server on :5000")
	if err := http.ListenAndServe(":5000", nil); err != nil {
		log.Fatal(err)
	}
}

func addData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	store.Lock()
	store.data[req.Key] = req.Value
	store.Unlock()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Data added successfully!")
}

func getData(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/data/"):]
	store.RLock()
	value, ok := store.data[key]
	store.RUnlock()
	if !ok {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{key: value})
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func backgroundTask() {
	for {
		store.RLock()
		fmt.Println("Current data store:", store.data)
		store.RUnlock()
		time.Sleep(10 * time.Second)
	}
}
