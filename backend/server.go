package main

import (
	"fmt"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type search struct {
	id string `json:"id"`
}

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Use the \"/youtube/\" on the url (https://go-video.herokuapp.com/youtube/) and past the YouTube video URL after that \n" + 
				  "or use \"/instagram/\" to download Instagram videos!")
}

// Download downloads file on browser
func Download(w http.ResponseWriter, r *http.Request, domain string) {
	params := mux.Vars(r)
	var id = params["id"]
	id = FormatMailiciousURL(id)
	url := domain + id

	DownloadVideo(url)

	videoPath := "video.mp4"
	data, err := ioutil.ReadFile(videoPath)
	if err != nil {
		log.Print(err)
		DeleteVideo()
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+videoPath)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	http.ServeContent(w, r, videoPath, time.Now(), bytes.NewReader(data))

	DeleteVideo()
}

// DownloadVideo downloads video from youtube
func DownloadVideo(url string) {
	log.Println("Received   | " + url)

	cmd := "youtube-dl " + url + " -o \"/video-downloader/video.mp4\""

	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Print(err)
		return
	}

	log.Println("Downloaded | " + url)
}

// DeleteVideo deletes video file
func DeleteVideo() {
	cmd := "rm *.mp4"
	exec.Command("sh", "-c", cmd).Output()
}

// FormatMailiciousURL formats a URL if it has malicious character
func FormatMailiciousURL(input string) string {
	formatedInput := strings.Replace(input, "&", "", -1)
	return formatedInput
}

// Routes

func Youtube(w http.ResponseWriter, r *http.Request) {
	domain := "https://www.youtube.com/watch?v="
	Download(w, r, domain)
}

func Instagram(w http.ResponseWriter, r *http.Request) {
	domain := "https://www.instagram.com/p/"
	Download(w, r, domain)
}


// Main function
func main() {
	log.Print("Go-Video started!")

	router := mux.NewRouter()
	port := os.Args[1]

	// Routes
	router.HandleFunc("/youtube/{id}", Youtube).Methods("GET")
	router.HandleFunc("/instagram/{id}", Instagram).Methods("GET")

	router.HandleFunc("/", Home).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
