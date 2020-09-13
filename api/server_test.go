package main

import (
	"os/exec"
	"testing"
)

// TestDeleteVideo unit test for deleteVideo()
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

// TestGetSupportedPlataforms unit test for getSupportedPlataforms()
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

// // TestValidateURL unit test for validateURL()
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
