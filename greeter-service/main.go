package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var greeting string = "Hello"
var delay int = 10

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	time.Sleep(time.Duration(delay) * time.Millisecond)

	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"greeting": fmt.Sprintf("%s %s", greeting, name),
	})
}

func main() {
	greetingVariable := os.Getenv("GREETING")

	if greeting != "" {
		log.Println("Loaded greeting from environment variable: " + greeting)
		greeting = greetingVariable
	}

	d := os.Getenv("DELAY")

	if n, err := strconv.Atoi(d); err == nil {
		log.Println("Loaded delay from environment variable: " + d)
		delay = n
	}

	router := mux.NewRouter();

	router.HandleFunc("/greeting/{name}", greetingHandler)

	log.Println("Starting server on port 8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}