package db

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetDB(t *testing.T) {
	mockDB := &sql.DB{}

	database := &Database{db: mockDB}

	result := database.GetDB()

	if result != mockDB {
		t.Errorf("Ожидалось, что результат GetDB() будет равен тестовой базе данных, но получен: %v", result)
	}
}

func TestDatabase_ClearLobbyTable(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	database := &Database{db: db}

	mock.ExpectExec("DELETE FROM lobby").WillReturnResult(sqlmock.NewResult(0, 0))

	err := database.ClearLobbyTable()
	if err != nil {
		t.Errorf("Ожидалось, что ClearLobbyTable() выполнится без ошибок, но получено: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Ожидалось, что все ожидаемые запросы будут выполнены, но получено: %s", err)
	}
}

func TestDatabase_ClearAccLobbyTable(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	database := &Database{db: db}

	mock.ExpectExec("DELETE FROM accLobby").WillReturnResult(sqlmock.NewResult(0, 0))

	err := database.ClearAccLobbyTable()
	if err != nil {
		t.Errorf("Ожидалось, что ClearAccLobbyTable() выполнится без ошибок, но получено: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Ожидалось, что все ожидаемые запросы будут выполнены, но получено: %s", err)
	}
}
