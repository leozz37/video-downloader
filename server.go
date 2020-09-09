package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)

// DownloadYoutubeVideo downloads video from youtube
func DownloadYoutubeVideo(id string) {
	url := "https://www.youtube.com/watch?v=" + id

	cmd := "youtube-dl " + url + " -o video.mp4"
	exec.Command("sh", "-c", cmd).Output()

	fmt.Println("Video " + url + " downloaded!")
}

// DeleteVideo deletes video file
func DeleteVideo() {
	cmd := "rm video.mp4"
	exec.Command("sh", "-c", cmd).Output()
}

// YoutubeDownload downloads file on browser
func YoutubeDownload(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]

	DownloadYoutubeVideo(id)

	videoPath := "video.mp4"
	data, err := ioutil.ReadFile(videoPath)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+videoPath)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	http.ServeContent(w, r, videoPath, time.Now(), bytes.NewReader(data))

	DeleteVideo()
}

func main() {
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/youtube/{id}", YoutubeDownload).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
