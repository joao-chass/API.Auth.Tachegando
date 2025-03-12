package tests

import (
	"AuthTachegando/internal/db"
	"AuthTachegando/internal/handlers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Carrega o banco de dados de teste
	os.Setenv("DB_NAME", os.Getenv("DB_NAME_TEST"))
	db.Init()
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	return r
}

func TestRegisterUser(t *testing.T) {
	router := setupRouter()

	// Simula requisição para registrar usuário
	body := []byte(`{"username":"testeuser", "password":"123456"}`)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestLoginUser(t *testing.T) {
	router := setupRouter()

	// Primeiro, registrar um usuário
	registerBody := []byte(`{"username":"testelogin", "password":"senha123"}`)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(registerBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Agora, tentar fazer login
	loginBody := []byte(`{"username":"testelogin", "password":"senha123"}`)
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(loginBody))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)
	_, exists := response["token"]
	assert.True(t, exists, "Token não foi gerado corretamente")
}
