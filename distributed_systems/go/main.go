package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var dataStore = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

func main() {
	http.HandleFunc("/data", addData)
	http.HandleFunc("/data/", getData)
	go backgroundTask()
	http.ListenAndServe(":5000", nil)
}

func addData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		dataStore.Lock()
		dataStore.m[key] = value
		dataStore.Unlock()
		fmt.Fprintf(w, "Data added successfully!")
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[len("/data/"):]
	dataStore.RLock()
	value, ok := dataStore.m[key]
	dataStore.RUnlock()
	if ok {
		fmt.Fprintf(w, value)
	} else {
		fmt.Fprintf(w, "Key not found!")
	}
}

func backgroundTask() {
	for {
		fmt.Println("Running background task...")
		time.Sleep(10 * time.Second)
	}
}
