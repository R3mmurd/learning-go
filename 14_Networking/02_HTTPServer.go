package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

var (
	db     map[string]interface{} = make(map[string]interface{})
	dbLock sync.Mutex
)

type Entry struct {
	Key   string      `json:"key"`
	Value interface{} `json:value`
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Everything ok!")
}

func sendResponse(entry *Entry, w http.ResponseWriter, s int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(s)
	enc := json.NewEncoder(w)
	if err := enc.Encode(entry); err != nil {
		log.Printf("error encoding %+v - %s", entry, err)
	}
}

func dbGetHandler(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Path[4:]

	dbLock.Lock()
	defer dbLock.Unlock()

	value, ok := db[k]

	if !ok {
		http.Error(w, fmt.Sprintf("Key %q not found", k), http.StatusNotFound)
		return
	}

	entry := &Entry{
		Key:   k,
		Value: value,
	}

	sendResponse(entry, w, 200)
}

func dbPostHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	entry := &Entry{}

	if err := dec.Decode(entry); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbLock.Lock()
	defer dbLock.Unlock()
	db[entry.Key] = entry.Value

	sendResponse(entry, w, 201)
}

func main() {
	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/db", dbPostHandler)
	http.HandleFunc("/db/", dbGetHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
