package notes

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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

	notes, err := List(db)
	assert.NoError(t, err)
	assert.Len(t, notes, 2)
	assert.Equal(t, notes[0].Title, "Test Note 1")
	assert.Equal(t, notes[0].Content, "This is a test note.")
	assert.Equal(t, notes[1].Title, "Test Note 2")
	assert.Equal(t, notes[1].Content, "Another test note.")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "content"}).
		AddRow(1, "Test Note", "This is a test note.")

	mock.ExpectQuery("SELECT \\* FROM notes where ID = \\$1").
		WithArgs("1").
		WillReturnRows(rows)

	note, err := Get(db, "1")
	assert.NoError(t, err)
	assert.Equal(t, note.Title, "Test Note")
	assert.Equal(t, note.Content, "This is a test note.")

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("DELETE FROM notes WHERE id = \\$1").
		WithArgs("1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Delete(db, "1")
	assert.NoError(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	note := Note{
		Title:   "Test Note",
		Content: "This is a test note.",
	}

	mock.ExpectQuery("INSERT INTO notes \\(title, content\\) VALUES \\(\\$1, \\$2\\) RETURNING id").
		WithArgs(note.Title, note.Content).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	err = Create(db, note)
	assert.NoError(t, err)
	assert.Equal(t, note.ID, 0)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	note := Note{
		Title:   "Updated Test Note",
		Content: "This is an updated test note.",
	}

	mock.ExpectExec("UPDATE notes SET title = \\$1, content = \\$2 WHERE id = \\$3").
		WithArgs(note.Title, note.Content, "1").
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = Update(db, "1", note)
	assert.NoError(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateNothing(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	note := Note{
		Title:   "Updated Test Note",
		Content: "This is an updated test note.",
	}

	mock.ExpectExec("UPDATE notes SET title = \\$1, content = \\$2 WHERE id = \\$3").
		WithArgs(note.Title, note.Content, "1").
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = Update(db, "1", note)
	assert.Error(t, err)

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestInitializeTable(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectExec("CREATE TABLE IF NOT EXISTS notes \\(id SERIAL PRIMARY KEY, title TEXT, content TEXT\\)").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = InitializeTable(db)
	assert.NoError(t, err)

	assert.NoError(t, mock.ExpectationsWereMet())
}
