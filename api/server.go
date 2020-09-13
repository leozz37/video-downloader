package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/rs/cors"
)

// APIRequest request struc
type APIRequest struct {
	URL string `json:"data"`
}

// download work as a main func
func download(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var payload APIRequest
	decoder.Decode(&payload)

	if !validateURL(payload.URL) {
		log.Println("Invalid URL")
		return
	}

	downloadVideo(payload.URL)

	videoPath := "video.mp4"
	data, err := ioutil.ReadFile(videoPath)
	if err != nil {
		log.Print(err)
		deleteVideo()
		return
	}

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename="+videoPath)
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	http.ServeContent(w, r, videoPath, time.Now(), bytes.NewReader(data))

	deleteVideo()
}

// downloadVideo uses youtube-dl to download videos
func downloadVideo(URL string) {

	log.Println("Received   | " + URL)

	cmd := "youtube-dl " + URL + " -o video.mp4"
	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Print(err)
		return
	}

	log.Println("Downloaded | " + URL)
}

// deleteVideo deletes video file
func deleteVideo() {

	cmd := "rm *.mp4"
	exec.Command("sh", "-c", cmd).Output()
}

// getSupportedPlataforms returns the supported video plataforms
func getSupportedPlataforms() [4]string {

	return [...]string{
		"facebook",
		"twitter",
		"youtube",
		"instagram",
	}
}

// validateURL check for a valid URL domain (youtube, twitter, instagram...)
func validateURL(URL string) bool {

	suportedPlataforms := getSupportedPlataforms()

	for _, plataform := range suportedPlataforms {

		if strings.Contains(URL, plataform) {
			return true
		}
	}

	return false
}

func main() {

	log.Print("Go-Video started!")

	mux := http.NewServeMux()
	mux.HandleFunc("/download", download)

	handler := cors.Default().Handler(mux)

	port := os.Args[1]
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
