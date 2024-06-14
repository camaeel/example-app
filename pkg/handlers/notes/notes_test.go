package notes

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	sqlmock "github.com/DATA-DOG/go-sqlmock"
// 	"github.com/camaeel/example-app/pkg/models/notes"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestList(t *testing.T) {
// 	// Create a mock database
// 	db, _, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	// Define test notes data
// 	testNotes := []notes.Note{
// 		{
// 			Title:   "Test Note 1",
// 			Content: "This is a test note.",
// 		},
// 		{
// 			Title:   "Test Note 2",
// 			Content: "Another test note.",
// 		},
// 	}

// 	// Create a test router
// 	router := gin.New()
// 	router.GET("/notes", func(c *gin.Context) {
// 		c.Set("db", db)
// 		List(c)
// 	})

// 	// Create a test request
// 	req := httptest.NewRequest("GET", "/notes", nil)

// 	// Create a test response recorder
// 	w := httptest.NewRecorder()

// 	// Serve the request
// 	router.ServeHTTP(w, req)

// 	// Assert the expected response
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `[{"title":"Test Note 1","body":"This is a test note."},{"title":"Test Note 2","body":"Another test note."}]`, w.Body.String())
// }

// func TestGet(t *testing.T) {
// 	// Create a mock database
// 	db := new(MockDB)

// 	// Define test note data
// 	testNote := notes.Note{
// 		Title:   "Test Note",
// 		Content: "This is a test note.",
// 	}

// 	// Mock database response
// 	db.On("Get", "1").Return(&testNote, nil)

// 	// Create a test router
// 	router := gin.New()
// 	router.GET("/notes/:id", func(c *gin.Context) {
// 		c.Set("db", db)
// 		Get(c)
// 	})

// 	// Create a test request
// 	req := httptest.NewRequest("GET", "/notes/1", nil)

// 	// Create a test response recorder
// 	w := httptest.NewRecorder()

// 	// Serve the request
// 	router.ServeHTTP(w, req)

// 	// Assert the expected response
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{"title":"Test Note","body":"This is a test note."}`, w.Body.String())
// }

// func TestDelete(t *testing.T) {
// 	// Create a mock database
// 	db := new(MockDB)

// 	// Mock database response
// 	db.On("Delete", "1").Return(nil)

// 	// Create a test router
// 	router := gin.New()
// 	router.DELETE("/notes/:id", func(c *gin.Context) {
// 		c.Set("db", db)
// 		Delete(c)
// 	})

// 	// Create a test request
// 	req := httptest.NewRequest("DELETE", "/notes/1", nil)

// 	// Create a test response recorder
// 	w := httptest.NewRecorder()

// 	// Serve the request
// 	router.ServeHTTP(w, req)

// 	// Assert the expected response
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{"result":"ok"}`, w.Body.String())
// }

// func TestCreate(t *testing.T) {
// 	// Create a mock database
// 	db := new(MockDB)

// 	// Define test note data
// 	testNote := notes.Note{
// 		Title:   "Test Note",
// 		Content: "This is a test note.",
// 	}

// 	// Mock database response
// 	db.On("Create", testNote).Return(nil)

// 	// Create a test router
// 	router := gin.New()
// 	router.POST("/notes", func(c *gin.Context) {
// 		c.Set("db", db)
// 		Create(c)
// 	})

// 	// Create a test request
// 	reqBody, _ := json.Marshal(testNote)
// 	req := httptest.NewRequest("POST", "/notes", bytes.NewBuffer(reqBody))

// 	// Create a test response recorder
// 	w := httptest.NewRecorder()

// 	// Serve the request
// 	router.ServeHTTP(w, req)

// 	// Assert the expected response
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{"result":"ok"}`, w.Body.String())
// }

// func TestUpdate(t *testing.T) {
// 	// Create a mock database
// 	db := new(MockDB)

// 	// Define test note data
// 	testNote := notes.Note{
// 		Title:   "Updated Test Note",
// 		Content: "This is an updated test note.",
// 	}

// 	// Mock database response
// 	db.On("Update", "1", testNote).Return(nil)

// 	// Create a test router
// 	router := gin.New()
// 	router.PUT("/notes/:id", func(c *gin.Context) {
// 		c.Set("db", db)
// 		Update(c)
// 	})

// 	// Create a test request
// 	reqBody, _ := json.Marshal(testNote)
// 	req := httptest.NewRequest("PUT", "/notes/1", bytes.NewBuffer(reqBody))

// 	// Create a test response recorder
// 	w := httptest.NewRecorder()

// 	// Serve the request
// 	router.ServeHTTP(w, req)

// 	// Assert the expected response
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.JSONEq(t, `{"result":"ok"}`, w.Body.String())
// }
