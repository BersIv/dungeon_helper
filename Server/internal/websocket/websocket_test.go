package websocket

import (
	"context"
	"database/sql"
	"dungeons_helper/internal/character"
	"dungeons_helper/utilMocks"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/websocket"
)

func TestCreateLobbyWebSocket(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO lobby`).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	handler := NewHandler(db, NewHub(db), &utilMocks.MockTokenGetter{Id: 123, Err: nil})
	s := httptest.NewServer(http.HandlerFunc(handler.CreateLobby))

	defer s.Close()

	u := "ws" + strings.TrimPrefix(s.URL, "http") + "/lobby/create?lobbyName=1&lobbyPassword=1&amount=2"
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("error connecting to websocket server: %v", err)
	}
	defer ws.Close()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestJoinLobbyWebSocket(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	mockRows := sqlmock.NewRows([]string{"id"}).
		AddRow(0)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT idChar FROM accChar WHERE idAccount = ? AND act = 1`)).WithArgs("0").WillReturnRows(mockRows)

	handler := NewHandler(db, NewHub(db), &utilMocks.MockTokenGetter{Id: 123, Err: nil})
	s := httptest.NewServer(http.HandlerFunc(handler.JoinLobby))

	defer s.Close()

	u := "ws" + strings.TrimPrefix(s.URL, "http") + "/lobby/join?idLobby=1"

	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("error connecting to websocket server: %v", err)
	}
	defer ws.Close()
}

type Database struct {
	db *sql.DB
}

func (d *Database) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return d.db.ExecContext(ctx, query, args...)
}

func (d *Database) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return d.db.PrepareContext(ctx, query)
}

func (d *Database) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return d.db.QueryContext(ctx, query, args...)
}

func (d *Database) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return d.db.QueryRowContext(ctx, query, args...)
}

func (d *Database) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return d.db.BeginTx(ctx, opts)
}

func TestHub_Run(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer mockDB.Close()

	dbInstance := &Database{db: mockDB}

	hub := NewHub(dbInstance)
	ctx := context.Background()
	go hub.Run()

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO accLobby(idAcc, idLobby) VALUES (?, ?)")).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	mock.ExpectBegin()
	rows := sqlmock.NewRows([]string{"lobbyMasterId"}).AddRow(1)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT lobbyMasterId FROM lobby WHERE id = ?")).WillReturnRows(rows)

	mock.ExpectExec("DELETE FROM accLobby WHERE idLobby = ?").WillReturnResult(sqlmock.NewResult(0, 1))

	testClient := &Client{
		Id:        1,
		IdLobby:   1,
		Nickname:  "TestClient",
		Context:   ctx,
		Character: &character.Character{},
	}

	fmt.Printf("t: %v\n", testClient)
	hub.JoinRoom <- testClient

	// if _, ok := hub.LobbyMembers[testClient.IdLobby]; !ok {
	// 	t.Errorf("Client not added to the lobby")
	// }

	//hub.LeaveRoom <- testClient

	// if _, ok := hub.LobbyMembers[testClient.IdLobby]; ok {
	// 	t.Errorf("Client not removed from the lobby")
	// }

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Not all expectations were met: %v", err)
	// }
}
