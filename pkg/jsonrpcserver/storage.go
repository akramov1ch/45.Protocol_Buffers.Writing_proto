package jsonrpcserver

import (
	"45HW/pkg/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() *Storage {
	db, err := sql.Open("postgres", config.Conf.DatabaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	createTableQuery := `
    CREATE TABLE IF NOT EXISTS words (
        id SERIAL PRIMARY KEY,
        word TEXT NOT NULL
    );`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	return &Storage{db: db}
}

func (s *Storage) Create(word string) error {
	insertQuery := `INSERT INTO words (word) VALUES ($1)`
	_, err := s.db.Exec(insertQuery, word)
	if err != nil {
		return fmt.Errorf("Error inserting word: %v", err)
	}
	return nil
}

func (s *Storage) Read(id int) (string, error) {
	var word string
	selectQuery := `SELECT word FROM words WHERE id = $1`
	err := s.db.QueryRow(selectQuery, id).Scan(&word)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("Word not found")
		}
		return "", fmt.Errorf("Error reading word: %v", err)
	}
	return word, nil
}

func (s *Storage) Update(id int, newWord string) error {
	updateQuery := `UPDATE words SET word = $1 WHERE id = $2`
	_, err := s.db.Exec(updateQuery, newWord, id)
	if err != nil {
		return fmt.Errorf("Error updating word: %v", err)
	}
	return nil
}

func (s *Storage) Delete(id int) error {
	deleteQuery := `DELETE FROM words WHERE id = $1`
	_, err := s.db.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("Error deleting word: %v", err)
	}
	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}
