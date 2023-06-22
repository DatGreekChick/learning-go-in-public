package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"example.com/api" // need this to avoid duplication of Album struct creation
)

var router = setupRouter()

func TestGetAlbumsRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAlbumIdRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/albums/3", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostAlbums(t *testing.T) {
	newAlbum := api.Album{
		ID:     "4",
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}

	var b bytes.Buffer
	err := json.NewEncoder(&b).Encode(newAlbum)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/albums", &b)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
