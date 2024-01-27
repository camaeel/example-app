package notes

import (
	"database/sql"
	"fmt"
	"log"
)

func InitializeTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS notes (id SERIAL PRIMARY KEY, title TEXT, content TEXT)")
	return err
}

type Note struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func List(db *sql.DB) ([]Note, error) {
	rows, err := db.Query("SELECT * FROM notes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	notes := []Note{}
	for rows.Next() {
		var u Note
		err = rows.Scan(&u.ID, &u.Title, &u.Content)
		if err != nil {
			return []Note{}, err
		}
		notes = append(notes, u)
	}
	if err := rows.Err(); err != nil {
		return []Note{}, err
	}

	return notes, nil
}

func Get(db *sql.DB, id string) (Note, error) {
	rows := db.QueryRow("SELECT * FROM notes where ID = $1", id)
	note := Note{}
	err := rows.Scan(&note.ID, &note.Title, &note.Content)
	return note, err
}

func Delete(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM notes WHERE id = $1", id)
	return err
}

func Create(db *sql.DB, note Note) error {
	err := db.QueryRow("INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id", note.Title, note.Content).Scan(&note.ID)
	return err
}

func Update(db *sql.DB, note Note) error {
	return fmt.Errorf("not implemented yet")
}
