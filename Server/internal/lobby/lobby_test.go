package lobby

import (
	"context"
	"database/sql"
	"dungeons_helper/utilMocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockRepo struct{}

func (m *MockRepo) GetAllLobby(ctx context.Context) ([]GetLobbyRes, error) {
	return []GetLobbyRes{
		{
			Id: 1, LobbyName: "Первое", Amount: 1, PlayersInLobby: 1},
		{
			Id: 2, LobbyName: "Второе", Amount: 2, PlayersInLobby: 2}}, nil
}

type MockRepoError struct{}

func (m *MockRepoError) GetAllLobby(ctx context.Context) ([]GetLobbyRes, error) {
	return nil, errors.New("mocked error")
}

type MockService struct{}

func (s *MockService) GetAllLobby(ctx context.Context) ([]GetLobbyRes, error) {
	return []GetLobbyRes{}, nil
}

type MockServiceError struct{}

func (s *MockServiceError) GetAllLobby(ctx context.Context) ([]GetLobbyRes, error) {
	return nil, errors.New("fake service error")
}

func TestGetAllLobbyRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "lobbyName", "amount", "count(ac.idAcc)"}).
		AddRow(1, "Первое", 1, 1).
		AddRow(2, "Второе", 2, 2)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT l.id, l.lobbyName, l.amount, count(ac.idAcc) FROM lobby l 
					LEFT JOIN accLobby ac on l.id = ac.idLobby 
					LEFT JOIN account a on ac.idAcc = a.id
					GROUP BY l.id, l.lobbyName, a.id`)).WillReturnRows(mockRows)

	repo := NewRepository(db)

	ctx := context.Background()
	lobby, err := repo.GetAllLobby(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(lobby) != 2 {
		t.Errorf("expected 2 alignments, got %d", len(lobby))
	}

	expected := []GetLobbyRes{
		{Id: 1, LobbyName: "Первое", Amount: 1, PlayersInLobby: 1},
		{Id: 2, LobbyName: "Второе", Amount: 2, PlayersInLobby: 2},
	}

	if !reflect.DeepEqual(lobby, expected) {
		t.Errorf("expected %+v, got %+v", expected, lobby)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllAlignmentsRepositoryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT l.id, l.lobbyName, l.amount, count(ac.idAcc) FROM lobby l 
					LEFT JOIN accLobby ac on l.id = ac.idLobby 
					LEFT JOIN account a on ac.idAcc = a.id
					GROUP BY l.id, l.lobbyName, a.id`)).WillReturnError(sql.ErrConnDone)

	repo := NewRepository(db)

	ctx := context.Background()
	_, err = repo.GetAllLobby(ctx)
	if err == nil {
		t.Error("expected an error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllClassesService(t *testing.T) {
	mockRepo := &MockRepo{}
	service := NewService(mockRepo)

	ctx := context.Background()
	classes, err := service.GetAllLobby(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expected := []GetLobbyRes{
		{
			Id: 1, LobbyName: "Первое", Amount: 1, PlayersInLobby: 1},
		{
			Id: 2, LobbyName: "Второе", Amount: 2, PlayersInLobby: 2}}
	if !reflect.DeepEqual(classes, expected) {
		t.Errorf("expected %+v, got %+v", expected, classes)
	}
}

func TestGetAllClassesErrorService(t *testing.T) {
	mockRepo := &MockRepoError{}
	service := NewService(mockRepo)

	ctx := context.Background()
	_, err := service.GetAllLobby(ctx)
	if err == nil {
		t.Error("expected an error, got nil")
	} else if err.Error() != "mocked error" {
		t.Errorf("expected mocked error, got %v", err)
	}
}

func TestGetAllClassesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllLobby(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}

	var lobby []GetLobbyRes
	if err := json.Unmarshal(rr.Body.Bytes(), &lobby); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}
}

func TestGetAllClassesHandler_Unauthorized(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 0, Err: errors.New("authorization header is missing")}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllLobby(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	expectedError := "authorization header is missing"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}

func TestGetAllClassesHandler_ServiceError(t *testing.T) {
	req, err := http.NewRequest("GET", "/classes", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer fake-token")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllLobby(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expectedError := "fake service error"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}
