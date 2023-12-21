package tests

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"testing"
	"time"
)

var serverAddr string

func TestMain(m *testing.M) {
	coverdir := "coverage"
	_ = os.Remove(coverdir)
	_ = os.Mkdir(coverdir, 0700)
	if err := os.Setenv("GOCOVERDIR", coverdir); err != nil {
		log.Fatal(err)
	}
	defer os.Unsetenv("GOCOVERDIR")

	if err := os.Setenv("PORT", "8089"); err != nil {
		log.Fatal(err)
	}
	defer os.Unsetenv("PORT")
	serverAddr = "http://localhost:8089"

	cmd := exec.Command("../server")
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// wait few secs for the server to start
	time.Sleep(time.Second * 3)

	m.Run()

	_, err := http.Get(serverAddr + "/exit")
	if err != nil {
		log.Fatal(err)
	}
}
