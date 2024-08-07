package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

type DatabaseTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

func NewDatabase() (*Database, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	database := &Database{db: db}

	err = database.ClearLobbyTable()
	if err != nil {
		return nil, err
	}

	err = database.ClearAccLobbyTable()
	if err != nil {
		return nil, err
	}

	return database, nil
}

func (d *Database) Close() {
	err := d.db.Close()
	if err != nil {
		return
	}
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func (d *Database) ClearLobbyTable() error {
	_, err := d.db.Exec("DELETE FROM lobby")
	return err
}

func (d *Database) ClearAccLobbyTable() error {
	_, err := d.db.Exec("DELETE FROM accLobby")
	return err
}
