package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// GetURL gets video URL from URL
func GetURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var url = params["id"]

	fmt.Println("https://www.youtube.com/watch?v=" + url)
}

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/download/{id}", GetURL).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
