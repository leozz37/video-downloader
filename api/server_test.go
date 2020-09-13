package main

import (
	"os/exec"
	"testing"
)

// TestDownloadVideo unit test for downloadVideo function
func TestDownloadVideo(t *testing.T) {

	URL := "https://www.youtube.com/watch?v=UO_QuXr521I"
	plataform := "youtube"

	downloadVideo(URL, plataform)

	cmd := "ls video.mp4"
	result, _ := exec.Command("sh", "-c", cmd).Output()

	if string(result) == "" {
		t.Errorf("Video was not downloaded")
	}

	deleteVideo()
}

// TestDeleteVideo unit test for deleteVideo function
func TestDeleteVideo(t *testing.T) {

	cmd := "touch video.mp4"
	exec.Command("sh", "-c", cmd).Output()

	deleteVideo()

	cmd = "ls video.mp4"
	result, _ := exec.Command("sh", "-c", cmd).Output()

	if string(result) != "" {
		t.Errorf("Video was not deleted")
	}

}

// TestGetSupportedPlataforms unit test for getSupportedPlataforms function
func TestGetSupportedPlataforms(t *testing.T) {

	plataforms := [4]string{
		"facebook",
		"twitter",
		"youtube",
		"instagram",
	}

	result := getSupportedPlataforms()

	for i := range result {
		if plataforms[i] != result[i] {
			t.Errorf("Plataform was incorrect, got: %s, want: %s.", result[i], plataforms[i])
		}
	}
}

// // TestValidateURL unit test for validateURL function
func TestValidateURL(t *testing.T) {

	plataforms := [4]string{
		"facebook",
		"twitter",
		"youtube",
		"instagram",
	}

	for _, plataform := range plataforms {

		URL := "www." + plataform + ".com/teste"

		result := validateURL(URL)

		if result != plataform {
			t.Errorf("Plataform was incorrect, got: %s, want: %s.", result, plataform)
		}
	}
}
