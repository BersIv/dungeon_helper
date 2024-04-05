package subraces

import (
	"bytes"
	"context"
	"database/sql"
	"dungeons_helper/internal/stats"
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

func (m *MockRepo) GetAllSubraces(ctx context.Context, race GetSubracesReq) ([]Subraces, error) {
	return []Subraces{{
		Id: 1, SubraceName: "Фул", Stats: stats.GetStatsRes{
			Strength:     1,
			Dexterity:    1,
			Constitution: 1,
			Intelligence: 1,
			Wisdom:       1,
			Charisma:     1,
		}},
		{
			Id: 2, SubraceName: "Полу", Stats: stats.GetStatsRes{
				Strength:     2,
				Dexterity:    2,
				Constitution: 2,
				Intelligence: 2,
				Wisdom:       2,
				Charisma:     2,
			}}}, nil
}

type MockService struct{}

func (s *MockService) GetAllSubraces(ctx context.Context, req GetSubracesReq) ([]Subraces, error) {
	return []Subraces{}, nil
}

type MockServiceError struct{}

func (s *MockServiceError) GetAllSubraces(ctx context.Context, req GetSubracesReq) ([]Subraces, error) {
	return nil, errors.New("fake service error")
}

type MockRepoError struct{}

func (m *MockRepoError) GetAllSubraces(ctx context.Context, req GetSubracesReq) ([]Subraces, error) {
	return nil, errors.New("mocked error")
}

func TestGetAllSubracesRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "subraceName", "strength", "dexterity", "constitution", "intelligence", "wisdom", "charisma"}).
		AddRow(1, "Фул", 1, 1, 1, 1, 1, 1).
		AddRow(2, "Полу", 2, 2, 2, 2, 2, 2)

	mock.ExpectQuery(`SELECT r.id, r.subraceName, s.strength, s.dexterity, s.constitution,
					s.intelligence, s.wisdom, s.charisma FROM subrace r
					LEFT JOIN stats s ON s.id = r.idStats
					WHERE idRace = ?`).WillReturnRows(mockRows)

	repo := NewRepository(db)

	ctx := context.Background()
	subraces, err := repo.GetAllSubraces(ctx, GetSubracesReq{IdRace: 1})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(subraces) != 2 {
		t.Errorf("expected 2 subraces, got %d", len(subraces))
	}

	expected := []Subraces{
		{Id: 1, SubraceName: "Фул", Stats: stats.GetStatsRes{
			Strength:     1,
			Dexterity:    1,
			Constitution: 1,
			Intelligence: 1,
			Wisdom:       1,
			Charisma:     1,
		}},
		{Id: 2, SubraceName: "Полу", Stats: stats.GetStatsRes{
			Strength:     2,
			Dexterity:    2,
			Constitution: 2,
			Intelligence: 2,
			Wisdom:       2,
			Charisma:     2,
		}},
	}

	if !reflect.DeepEqual(subraces, expected) {
		t.Errorf("expected %+v, got %+v", expected, subraces)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllSubracesRepositoryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT r.id, r.subraceName, s.strength, s.dexterity, s.constitution,
						s.intelligence, s.wisdom, s.charisma FROM subrace r
						LEFT JOIN stats s ON s.id = r.idStats
						WHERE idRace = ?`).WillReturnError(sql.ErrConnDone)

	repo := NewRepository(db)

	ctx := context.Background()
	_, err = repo.GetAllSubraces(ctx, GetSubracesReq{IdRace: 0})
	if err == nil {
		t.Error("expected an error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllSubracesService(t *testing.T) {
	mockRepo := &MockRepo{}
	service := NewService(mockRepo)

	ctx := context.Background()
	subraces, err := service.GetAllSubraces(ctx, GetSubracesReq{IdRace: 1})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expected := []Subraces{
		{Id: 1, SubraceName: "Фул", Stats: stats.GetStatsRes{
			Strength:     1,
			Dexterity:    1,
			Constitution: 1,
			Intelligence: 1,
			Wisdom:       1,
			Charisma:     1,
		}},
		{Id: 2, SubraceName: "Полу", Stats: stats.GetStatsRes{
			Strength:     2,
			Dexterity:    2,
			Constitution: 2,
			Intelligence: 2,
			Wisdom:       2,
			Charisma:     2,
		}},
	}
	if !reflect.DeepEqual(subraces, expected) {
		t.Errorf("expected %+v, got %+v", expected, subraces)
	}
}

func TestGetSubracesErrorService(t *testing.T) {
	mockRepo := &MockRepoError{}
	service := NewService(mockRepo)

	ctx := context.Background()
	_, err := service.GetAllSubraces(ctx, GetSubracesReq{IdRace: 1})
	if err == nil {
		t.Error("expected an error, got nil")
	} else if err.Error() != "mocked error" {
		t.Errorf("expected mocked error, got %v", err)
	}
}

func TestGetAllSubracesHandler(t *testing.T) {
	requestBody := []byte(`{"idRace": 123}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()
	handler.GetAllSubraces(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}

	var subraces []Subraces
	if err := json.Unmarshal(rr.Body.Bytes(), &subraces); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}
}

func TestGetAllSubracesHandler_Unauthorized(t *testing.T) {
	requestBody := []byte(`{"idRace": 123}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 0, Err: errors.New("authorization header is missing")}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllSubraces(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	expectedError := "authorization header is missing"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}

func TestGetAllSubracesHandler_ServiceError(t *testing.T) {
	requestBody := []byte(`{"idRace": 123}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer fake-token")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllSubraces(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expectedError := "fake service error"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}

func TestGetAllSubracesHandler_IdZeroError(t *testing.T) {
	requestBody := []byte(`{"idRace": 0}`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer fake-token")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllSubraces(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestGetAllSubracesHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{"idRace": 0`)

	req, err := http.NewRequest("GET", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer fake-token")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllSubraces(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
