package class

import (
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

func (m *MockRepo) GetAllClasses(ctx context.Context) ([]Class, error) {
	return []Class{{Id: 1, ClassName: "Воин"}, {Id: 2, ClassName: "Бард"}}, nil
}

type MockClassService struct{}

func (s *MockClassService) GetAllClasses(ctx context.Context) ([]Class, error) {
	return []Class{}, nil
}

type MockServiceError struct{}

func (s *MockServiceError) GetAllClasses(ctx context.Context) ([]Class, error) {
	return nil, errors.New("fake service error")
}

func TestGetAllClassesRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "className"}).
		AddRow(1, "Воин").
		AddRow(2, "Бард")

	mock.ExpectQuery("SELECT id, className FROM class").WillReturnRows(mockRows)

	repo := NewRepository(db)

	ctx := context.Background()
	classes, err := repo.GetAllClasses(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(classes) != 2 {
		t.Errorf("expected 2 classes, got %d", len(classes))
	}

	expected := []Class{
		{Id: 1, ClassName: "Воин"},
		{Id: 2, ClassName: "Бард"},
	}

	for i, c := range classes {
		if c.Id != expected[i].Id || c.ClassName != expected[i].ClassName {
			t.Errorf("expected %+v, got %+v", expected[i], c)
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllClassesRepository_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, className FROM class").WillReturnError(sql.ErrConnDone)

	repo := NewRepository(db)

	ctx := context.Background()
	_, err = repo.GetAllClasses(ctx)
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
	classes, err := service.GetAllClasses(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expected := []Class{{Id: 1, ClassName: "Воин"}, {Id: 2, ClassName: "Бард"}}
	if !reflect.DeepEqual(classes, expected) {
		t.Errorf("expected %+v, got %+v", expected, classes)
	}
}

type MockRepoError struct{}

func (m *MockRepoError) GetAllClasses(ctx context.Context) ([]Class, error) {
	return nil, errors.New("mocked error")
}

func TestGetAllClassesErrorService(t *testing.T) {
	mockRepo := &MockRepoError{}
	service := NewService(mockRepo)

	ctx := context.Background()
	_, err := service.GetAllClasses(ctx)
	if err == nil {
		t.Error("expected an error, got nil")
	} else if err.Error() != "mocked error" {
		t.Errorf("expected mocked error, got %v", err)
	}
}

func TestGetAllClassesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/classes", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockClassService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllClasses(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}

	var classes []Class
	if err := json.Unmarshal(rr.Body.Bytes(), &classes); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}
}

func TestGetAllClassesHandler_Unauthorized(t *testing.T) {
	req, err := http.NewRequest("GET", "/classes", nil)
	if err != nil {
		t.Fatal(err)
	}

	svc := &MockClassService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 0, Err: errors.New("authorization header is missing")}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllClasses(rr, req)

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

	handler.GetAllClasses(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expectedError := "fake service error"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}
