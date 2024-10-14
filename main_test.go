package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPingRoute(t *testing.T) {

	router := setupRouter()
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

// Test for POST /user/add

func TestPostUser(t *testing.T) {

	router := setupRouter()

	router = postUser(router)
	w := httptest.NewRecorder()

	exampleUser := User{
		Username: "test_name",
		Gender:   "male",
	}

	userJson, _ := json.Marshal(exampleUser)

	req, _ := http.NewRequest("POST", "/user/add", strings.NewReader(string(userJson)))

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	assert.Equal(t, string(userJson), w.Body.String())

}
