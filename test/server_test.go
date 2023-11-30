package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBook(t *testing.T) {
	res, err := http.Post("http://localhost:8088", "application/json", strings.NewReader(`{"name":"foo"}`))
	require.NoError(t, err)

	m := map[string]any{}
	err = json.NewDecoder(res.Body).Decode(&m)
	assert.NoError(t, err)

	fmt.Println(m)
}
