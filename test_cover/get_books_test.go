package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetBooks(t *testing.T) {
	res, err := http.Get(serverAddr + "/books")
	require.NoError(t, err)

	books := []map[string]any{}
	err = json.NewDecoder(res.Body).Decode(&books)
	assert.NoError(t, err)

	assert.Len(t, books, 10)
}

func TestGetBooksLimit(t *testing.T) {
	tt := []struct {
		name       string
		limit      string
		want       int
		statusCode int
	}{
		{name: "lower limit", limit: "4", want: 4, statusCode: http.StatusOK},
		{name: "higher limit", limit: "13", want: 13, statusCode: http.StatusOK},
		{name: "too high limit", limit: "45", want: 10, statusCode: http.StatusOK},
		{name: "negative limit", limit: "-10", want: 10, statusCode: http.StatusOK},
		{name: "invalid limit", limit: "djdj", statusCode: http.StatusBadRequest},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			endpoint := fmt.Sprintf("%s/books?limit=%s", serverAddr, tc.limit)
			res, err := http.Get(endpoint)

			assert.NoError(t, err)
			assert.Equal(t, tc.statusCode, res.StatusCode)

			if tc.statusCode == http.StatusOK {
				books := []map[string]any{}
				err = json.NewDecoder(res.Body).Decode(&books)

				assert.NoError(t, err)
				assert.Len(t, books, tc.want)
			}
		})
	}
}
