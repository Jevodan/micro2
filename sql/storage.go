package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Storage struct {
	db *sql.DB
}

func InitDb() *Storage {
	db, err := sql.Open("mysql", "root:serpent@/tasks")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Соединение с БД")
	}
	storage := Storage{
		db: db,
	}
	return &storage
}

func (s *Storage) GetName(id int) (string, error) {
	var name string
	err := s.db.QueryRow("SELECT name FROM users WHERE id= ?", id).Scan(&name)

	if err != nil {
		log.Fatal(err)
	}
	return name, nil
}
