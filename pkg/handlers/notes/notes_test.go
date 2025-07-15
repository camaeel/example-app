package notes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/camaeel/example-app/pkg/models/notes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "content"}).
		AddRow(1, "Test Note 1", "This is a test note.").
		AddRow(2, "Test Note 2", "Another test note.")

	mock.ExpectQuery("SELECT \\* FROM notes").WillReturnRows(rows)

	router := gin.New()
	router.GET("/notes", func(c *gin.Context) {
		c.Set("db", db)
		List(c)
	})

	req := httptest.NewRequest(http.MethodGet, "/notes", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var notes []notes.Note
	err = json.Unmarshal(w.Body.Bytes(), &notes)
	assert.NoError(t, err)
	assert.Len(t, notes, 2)
	assert.Equal(t, notes[0].Title, "Test Note 1")
	assert.Equal(t, notes[0].Content, "This is a test note.")
	assert.Equal(t, notes[1].Title, "Test Note 2")
	assert.Equal(t, notes[1].Content, "Another test note.")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "content"}).
		AddRow(1, "Test Note", "This is a test note.")

	mock.ExpectQuery("SELECT \\* FROM notes where ID = \\$1").
		WithArgs("1").
		WillReturnRows(rows)

	router := gin.New()
	router.GET("/notes/:id", func(c *gin.Context) {
		c.Set("db", db)
		Get(c)
	})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("db", db)

	req := httptest.NewRequest(http.MethodGet, "/notes/1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var note notes.Note
	err = json.Unmarshal(w.Body.Bytes(), &note)
	assert.NoError(t, err)
	assert.Equal(t, "Test Note", note.Title)
	assert.Equal(t, "This is a test note.", note.Content)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetNotExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "content"})

	mock.ExpectQuery("SELECT \\* FROM notes where ID = \\$1").
		WithArgs("1").
		WillReturnRows(rows)

	router := gin.New()
	router.GET("/notes/:id", func(c *gin.Context) {
		c.Set("db", db)
		Get(c)
	})

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Set("db", db)

	req := httptest.NewRequest(http.MethodGet, "/notes/1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Empty(t, w.Body.Bytes())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("DELETE FROM notes WHERE id = \\$1").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	router := gin.New()
	router.DELETE("/notes/:id", func(c *gin.Context) {
		c.Set("db", db)
		Delete(c)
	})

	req := httptest.NewRequest(http.MethodDelete, "/notes/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, response["result"], "ok")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	note := notes.Note{
		Title:   "Test Note",
		Content: "This is a test note.",
	}

	mock.ExpectQuery("INSERT INTO notes \\(title, content\\) VALUES \\(\\$1, \\$2\\) RETURNING id").
		WithArgs(note.Title, note.Content).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	reqBody, err := json.Marshal(note)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	ctx, engine := gin.CreateTestContext(w)
	engine.POST("/notes", func(c *gin.Context) {
		c.Set("db", db)
		Create(c)
	})
	ctx.Set("db", db)

	req := httptest.NewRequest(http.MethodPost, "/notes", bytes.NewBuffer(reqBody))
	engine.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, response["result"], "ok")

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	note := notes.Note{
		Title:   "Updated Test Note",
		Content: "This is an updated test note.",
	}

	mock.ExpectExec("UPDATE notes SET title = \\$1, content = \\$2 WHERE id = \\$3").
		WithArgs(note.Title, note.Content, "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	router := gin.New()
	router.PUT("/notes/:id", func(c *gin.Context) {
		c.Set("db", db)
		Update(c)
	})

	reqBody, _ := json.Marshal(note)
	req := httptest.NewRequest(http.MethodPut, "/notes/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, response["result"], "ok")

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateNoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	note := notes.Note{
		Title:   "Updated Test Note",
		Content: "This is an updated test note.",
	}

	mock.ExpectExec("UPDATE notes SET title = \\$1, content = \\$2 WHERE id = \\$3").
		WithArgs(note.Title, note.Content, "1").
		WillReturnResult(sqlmock.NewResult(0, 0))

	router := gin.New()
	router.PUT("/notes/:id", func(c *gin.Context) {
		c.Set("db", db)
		Update(c)
	})

	reqBody, _ := json.Marshal(note)
	req := httptest.NewRequest(http.MethodPut, "/notes/1", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// var response map[string]interface{}
	// err = json.Unmarshal(w.Body.Bytes(), &response)
	// assert.NoError(t, err)
	// assert.Equal(t, response["result"], "ok")

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
