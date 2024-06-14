package middleware

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	w := httptest.NewRecorder()
	_, engine := gin.CreateTestContext(w)

	engine.Use(Logger())

	// Create a test request
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	engine.ServeHTTP(w, req)
}

func TestInsertDB(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	middleware := InsertDB(func() (*sql.DB, error) {
		return db, nil
	})

	middleware(ctx)

	// Assert that the database was set in the context
	result, exists := ctx.Get("db")
	assert.True(t, exists)
	assert.NotNil(t, result)
}
