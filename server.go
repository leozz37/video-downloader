package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/leozz37/video-downloader/youtube"
)

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Use the \"/youtube/\" on the url (https://go-video.herokuapp.com/youtube/) and past the YouTube video URL after that")
}

func main() {
	router := mux.NewRouter()
	port := os.Args[1]

	// Routes
	router.HandleFunc("/youtube/{id}", youtube.Download).Methods("GET")
	router.HandleFunc("/", Home).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
