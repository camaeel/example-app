package health

import (
	"github.com/DATA-DOG/go-sqlmock"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthz(t *testing.T) {
	// Create a test router
	router := gin.New()
	router.GET("/healthz", Healthz)

	// Create a test request
	req := httptest.NewRequest("GET", "/healthz", nil)

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the expected response
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message":"ok"}`, w.Body.String())
}

func TestReadyz(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("SELECT 1 FROM notes").WithoutArgs().WillReturnResult(sqlmock.NewResult(0, 0))
	// Create a test router
	router := gin.New()
	router.GET("/readyz", func(c *gin.Context) {
		c.Set("db", db)
		Readyz(c)
	})

	// Create a test request
	req := httptest.NewRequest("GET", "/readyz", nil)

	// Create a test response recorder
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// Assert the expected response
	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"message":"ok"}`, w.Body.String())
}
