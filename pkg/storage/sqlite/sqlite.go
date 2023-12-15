package sqlite_storage

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/artyomtugaryov/vpnbot/pkg/storage"
	"github.com/artyomtugaryov/vpnbot/pkg/utils"
	"github.com/artyomtugaryov/vpnbot/pkg/utils/errors"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStorage struct {
	database *sql.DB
}

const (
	filename = "database.db"
)

func New(basePath string) SQLiteStorage {

	if err := utils.MkDir(basePath); err != nil {
		log.Fatal(err)
	}

	databasePath := filepath.Join(basePath, filename)
	database, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}

	storage := SQLiteStorage{
		database: database,
	}

	if err := storage.Initialize(); err != nil {
		log.Fatal(err)
	}
	return storage
}

func (s *SQLiteStorage) Initialize() error {
	sqlStmt := `
	create table if not exists customers (id integer not null primary key, username text);
	create table if not exists messages (id integer not null primary key, messsage text);
	`
	if _, err := s.database.Exec(sqlStmt); err != nil {
		return errors.Wrap("Cannot initialize database", err)
	}
	return nil
}

func (s *SQLiteStorage) Close() error {
	return s.database.Close()
}

func (s *SQLiteStorage) SaveCusromer(customer *storage.Customer) error {
	tx, err := s.database.Begin()
	if err != nil {
		return errors.Wrap("Cannot save a customer", err)
	}
	sqlStmt := `
	insert into customers(username) values(?);
	`
	stmt, err := tx.Prepare(sqlStmt)
	if err != nil {
		return errors.Wrap("Cannot save a customer", err)
	}

	defer stmt.Close()

	if _, err = stmt.Exec(customer.Username); err != nil {
		return errors.Wrap("Cannot save a customer", err)
	}
	if err = tx.Commit(); err != nil {
		return errors.Wrap("Cannot save a customer", err)
	}
	return nil
}
