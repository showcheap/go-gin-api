package tests

import (
	"go-gin-api/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := routes.NewRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Pong!", w.Body.String())
}

// func TestLoginRoute(t *testing.T) {
// 	router := routes.NewRouter()
// 	c, _ := db.Init()
// 	defer c.Close()

// 	var loginForm struct {
// 		Email    string
// 		Password string
// 	}

// 	loginForm.Email = "admin@localhost.com"
// 	loginForm.Password = "admin"

// 	data, _ := json.Marshal(loginForm)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBufferString(string(data)))
// 	router.ServeHTTP(w, req)

// 	assert.Equal(t, 200, w.Code)
// 	// assert.Equal(t, "Pong!", w.Body.String())
// }
