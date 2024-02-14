package stats

import (
	"bytes"
	"context"
	"database/sql"
	"dungeons_helper/utilMocks"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockRepo struct{}

func (m *MockRepo) GetStatsById(ctx context.Context, id int64) (*Stats, error) {
	return &Stats{
		Strength:     1,
		Dexterity:    1,
		Constitution: 1,
		Intelligence: 1,
		Wisdom:       1,
		Charisma:     1,
	}, nil
}

type MockRepoError struct{}

func (m *MockRepoError) GetStatsById(ctx context.Context, id int64) (*Stats, error) {
	return nil, errors.New("mocked error")
}

type MockService struct{}

func (s *MockService) GetStatsById(ctx context.Context, id int64) (*GetStatsRes, error) {
	return &GetStatsRes{}, nil
}

type MockServiceError struct{}

func (s *MockServiceError) GetStatsById(ctx context.Context, id int64) (*GetStatsRes, error) {
	return nil, errors.New("fake service error")
}

func TestGetStatsByIdRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma"}).
		AddRow(1, 1, 1, 1, 1, 1)

	mock.ExpectQuery(`SELECT strength, dexterity, constitution, 
					intelligence, wisdom, charisma FROM stats 
					WHERE id = ?`).WillReturnRows(mockRows)

	repo := NewRepository(db)

	ctx := context.Background()
	stats, err := repo.GetStatsById(ctx, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	expected := &Stats{
		Strength:     1,
		Dexterity:    1,
		Constitution: 1,
		Intelligence: 1,
		Wisdom:       1,
		Charisma:     1,
	}

	if !reflect.DeepEqual(stats, expected) {
		t.Errorf("expected %+v, got %+v", expected, stats)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetStatsByIdRepositoryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT strength, dexterity, constitution, 
						intelligence, wisdom, charisma FROM stats 
						WHERE id = ?`).WillReturnError(sql.ErrConnDone)

	repo := NewRepository(db)

	ctx := context.Background()
	_, err = repo.GetStatsById(ctx, 1)
	if err == nil {
		t.Error("expected an error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetStatsByIdService(t *testing.T) {
	mockRepo := &MockRepo{}
	service := NewService(mockRepo)

	ctx := context.Background()
	stats, err := service.GetStatsById(ctx, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expected := &GetStatsRes{
		Strength:     1,
		Dexterity:    1,
		Constitution: 1,
		Intelligence: 1,
		Wisdom:       1,
		Charisma:     1,
	}

	if !reflect.DeepEqual(stats, expected) {
		t.Errorf("expected %+v, got %+v", expected, stats)
	}
}

func TestGetStatsByIdServiceError(t *testing.T) {
	mockRepo := &MockRepoError{}
	service := NewService(mockRepo)

	ctx := context.Background()
	_, err := service.GetStatsById(ctx, 1)
	if err == nil {
		t.Error("expected an error, got nil")
	} else if err.Error() != "mocked error" {
		t.Errorf("expected mocked error, got %v", err)
	}
}

func TestGetStatsByIdHandler(t *testing.T) {
	requestBody := []byte(`{"id": 123}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()
	handler.GetStatsById(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}

	var stats *GetStatsRes
	if err := json.Unmarshal(rr.Body.Bytes(), &stats); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}
}

func TestGetStatsByIdHandler_Unauthorized(t *testing.T) {
	requestBody := []byte(`{"id": 123}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 0, Err: errors.New("authorization header is missing")}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetStatsById(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	expectedError := "authorization header is missing"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}

func TestGetStatsByIdHandler_ServiceError(t *testing.T) {
	requestBody := []byte(`{"id": 123}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer fake-token")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetStatsById(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expectedError := "fake service error"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}

func TestGetStatsByIdHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer fake-token")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetStatsById(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestGetStatsByIdHandler_ZeroId(t *testing.T) {
	requestBody := []byte(`{"Id": 0}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer fake-token")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetStatsById(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
