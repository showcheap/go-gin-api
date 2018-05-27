package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-gin-api/db"
	"go-gin-api/routes"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/stretchr/testify/assert"
)

func TestLoginRoute(t *testing.T) {

	c, err := db.Init()

	if err != nil {
		fmt.Println("Connect database error: ", err)
	}

	defer c.Close()

	router := routes.NewRouter()

	var loginForm struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	loginForm.Email = "admin@localhost.com"
	loginForm.Password = "admin"

	data, _ := json.Marshal(loginForm)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBufferString(string(data)))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
