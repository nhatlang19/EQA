package service_test

import (
	"EQA/backend/app/repository"
	"EQA/backend/app/service"
	"EQA/backend/config"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setupAuthTest(db *gorm.DB) *service.AuthServiceImpl {
	os.Setenv("SECRET_KEY", "AVp1g0#8TDM3/4WT%xM4)6?;jF$8(v")
	os.Setenv("PORT", "8080")
	userRepositoryImpl := repository.UserRepositoryInit(db)
	authServiceImpl := service.AuthServiceInit(userRepositoryImpl)
	setupAuthTestDB(db)
	return authServiceImpl
}

func setupAuthTestDB(db *gorm.DB) {
	db.Exec("INSERT INTO users (username, password, created_at, updated_at, deleted_at, active) VALUES('admin', '$2a$15$H7jcIf5zk5qcMfMi6uqC/.VgfkYMwq6r4OyryOOEz2/0tZ/I.XvF.', CURRENT_TIMESTAMP, NULL, NULL, true)")
}

func teardownAuthTestDB(db *gorm.DB) {
	db.Exec("TRUNCATE TABLE users")

}
func teardownAuthTest(db *gorm.DB) {
	teardownAuthTestDB(db)
	os.Unsetenv("ENV")
	os.Unsetenv("SECRET_KEY")
	os.Unsetenv("PORT")
}

func TestLoginHandler(t *testing.T) {
	os.Setenv("ENV", "TEST")
	db := config.ConnectToDB()
	s := setupAuthTest(db)
	defer teardownAuthTest(db)

	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/api/auth/login", s.Login)

	t.Run("ValidRequest", func(t *testing.T) {
		reqBody := `{"username": "admin", "password": "123456"}`
		req, _ := http.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(reqBody))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
		assert.Contains(t, w.Body.String(), "accessToken")
	})

	t.Run("InvalidUsername", func(t *testing.T) {
		reqBody := `{"username": "testuser", "password": "123456"}`
		req, _ := http.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(reqBody))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusNotFound)
		assert.Contains(t, w.Body.String(), "User does not exist")
	})

	t.Run("InvalidPassword", func(t *testing.T) {
		reqBody := `{"username": "admin", "password": "1234567"}`
		w := httptest.NewRecorder()

		req, _ := http.NewRequest(http.MethodPost, "/api/auth/login", strings.NewReader(reqBody))
		router.ServeHTTP(w, req)

		assert.Equal(t, w.Code, http.StatusUnauthorized)
		assert.Contains(t, w.Body.String(), "Invalid username or password")
	})

}
