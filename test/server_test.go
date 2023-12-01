package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("main")

	cmd := exec.Command("../coverage")
	fmt.Println(cmd.Start())
	time.Sleep(time.Second * 5)

	m.Run()

	res, err := http.Get("http://localhost:8088/exit")
	fmt.Println(res, err)
}

func TestGetBooks(t *testing.T) {
	fmt.Println("TestGetBooks")

	res, err := http.Get("http://localhost:8088/book")
	require.NoError(t, err)

	books := []map[string]any{}
	err = json.NewDecoder(res.Body).Decode(&books)
	assert.NoError(t, err)

	fmt.Println(books)
}
