package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davidzech/webtutorial/notes"
	"github.com/stretchr/testify/require"
)

func TestNotes(t *testing.T) { // Test Suite is called TestNotes
	t.FailNow()
	router := createRouter()
	t.Run("create", func(t *testing.T) { // one test case
		recorder := httptest.NewRecorder()
		note := notes.Note{
			Value: "Foobar",
		}
		j, _ := json.Marshal(&note)
		body := bytes.NewBuffer(j)
		request := httptest.NewRequest("POST", "http://localhost/notes", body)

		router.ServeHTTP(recorder, request)

		require.Equal(t, http.StatusCreated, recorder.Code)
		require.NoError(t, json.Unmarshal([]byte(recorder.Body.String()), &note))

		t.Run("read", func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", fmt.Sprintf("http://localhost/notes/%d", note.ID), nil)
			router.ServeHTTP(recorder, request)
			handleRead(recorder, request)
			require.Equal(t, http.StatusOK, recorder.Code)
		})

	})

}
