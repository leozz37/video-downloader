package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// DownloadYoutubeVideo downloads video from youtube
func DownloadYoutubeVideo(id string) {
	id = FormatMailiciousURL(id)

	url := "https://www.youtube.com/watch?v=" + id
	cmd := "youtube-dl " + url + " -o video.mp4"

	exec.Command("sh", "-c", cmd).Output()
	fmt.Println("Video " + url + " downloaded!")
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

// YoutubeDownload downloads file on browser
func YoutubeDownload(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]

	DownloadYoutubeVideo(id)

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

// Home page
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Use the \"/youtube/\" on the url (https://go-video.herokuapp.com/youtube/) and past the YouTube video URL after that")
}

func main() {
	router := mux.NewRouter()
	port := os.Args[1]

	// Routes
	router.HandleFunc("/youtube/{id}", YoutubeDownload).Methods("GET")
	router.HandleFunc("/", Home).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
