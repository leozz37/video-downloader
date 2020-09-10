package instagram

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Download downloads file on browser
func Download(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]

	DownloadInstagramVideo(id)

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

// DownloadInstagramVideo downloads video from instagram
func DownloadInstagramVideo(id string) {
	id = FormatMailiciousURL(id)
	url := "https://www.instagram.com/p/" + id + "/"

	log.Println("INSTAGRAM | received " + url)
	cmd := "instalooter post \"" + url + "\" -T video -V"

	_, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("INSTAGRAM | video " + url + " downloaded!")
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
