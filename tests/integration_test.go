package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spike510/task-manager/internal/db"
	taskHttp "github.com/spike510/task-manager/internal/http"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter(t *testing.T) *gin.Engine {
	database, err := db.Connect()
	if err != nil {
		t.Fatal(err)
	}
	return taskHttp.NewRouter(database)
}

func TestIntegration_RegisterLoginAndTaskFlow(t *testing.T) {
	router := setupRouter(t)

	// 1. Rejestracja
	w := httptest.NewRecorder()
	body := `{"email":"test@user.com", "password":"secret123"}`
	req, _ := http.NewRequest("POST", "/users/register", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	// 2. Logowanie
	w = httptest.NewRecorder()
	body = `{"email":"test@user.com", "password":"secret123"}`
	req, _ = http.NewRequest("POST", "/users/login", bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var loginResp map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &loginResp)
	if err != nil {
		t.Fatal(err)
	}

	token := loginResp["token"]
	assert.NotEmpty(t, token)

	// 3. Utworzenie zadania
	w = httptest.NewRecorder()
	body = `{"title":"Test Task", "description":"Integration test"}`
	req, _ = http.NewRequest("POST", "/tasks", bytes.NewBufferString(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)

	// 4. Pobranie listy zadań
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/tasks", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
