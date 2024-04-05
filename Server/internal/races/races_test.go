package races

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

func (m *MockRepo) GetAllRaces(ctx context.Context) ([]Races, error) {
	return []Races{
		{
			Id: 1, RaceName: "Эльф"},
		{
			Id: 2, RaceName: "Дворф"}}, nil
}

type MockRepoError struct{}

func (m *MockRepoError) GetAllRaces(ctx context.Context) ([]Races, error) {
	return nil, errors.New("mocked error")
}

type MockService struct{}

func (s *MockService) GetAllRaces(ctx context.Context) ([]Races, error) {
	return []Races{}, nil
}

type MockServiceError struct{}

func (s *MockServiceError) GetAllRaces(ctx context.Context) ([]Races, error) {
	return nil, errors.New("fake service error")
}

func TestGetAllRacesRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "raceName"}).
		AddRow(1, "Эльф").
		AddRow(2, "Дворф")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, raceName FROM races`)).WillReturnRows(mockRows)

	repo := NewRepository(db)

	ctx := context.Background()
	races, err := repo.GetAllRaces(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(races) != 2 {
		t.Errorf("expected 2 alignments, got %d", len(races))
	}

	expected := []Races{
		{Id: 1, RaceName: "Эльф"},
		{Id: 2, RaceName: "Дворф"},
	}

	if !reflect.DeepEqual(races, expected) {
		t.Errorf("expected %+v, got %+v", expected, races)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllRacesRepositoryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, raceName FROM races`)).WillReturnError(sql.ErrConnDone)

	repo := NewRepository(db)

	ctx := context.Background()
	_, err = repo.GetAllRaces(ctx)
	if err == nil {
		t.Error("expected an error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllRacesService(t *testing.T) {
	mockRepo := &MockRepo{}
	service := NewService(mockRepo)

	ctx := context.Background()
	races, err := service.GetAllRaces(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expected := []Races{
		{
			Id: 1, RaceName: "Эльф"},
		{
			Id: 2, RaceName: "Дворф"}}
	if !reflect.DeepEqual(races, expected) {
		t.Errorf("expected %+v, got %+v", expected, races)
	}
}

func TestGetAllRacesErrorService(t *testing.T) {
	mockRepo := &MockRepoError{}
	service := NewService(mockRepo)

	ctx := context.Background()
	_, err := service.GetAllRaces(ctx)
	if err == nil {
		t.Error("expected an error, got nil")
	} else if err.Error() != "mocked error" {
		t.Errorf("expected mocked error, got %v", err)
	}
}

func TestGetAllRacesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllRaces(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}

	var races []Races
	if err := json.Unmarshal(rr.Body.Bytes(), &races); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}
}

func TestGetAllRacesHandler_Unauthorized(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 0, Err: errors.New("authorization header is missing")}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllRaces(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	expectedError := "authorization header is missing"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}

func TestGetAllRacesHandler_ServiceError(t *testing.T) {
	req, err := http.NewRequest("GET", "/classes", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer fake-token")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllRaces(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expectedError := "fake service error"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}
