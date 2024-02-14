package skills

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

func (m *MockRepo) GetAllSkills(ctx context.Context) ([]Skills, error) {
	return []Skills{
		{
			Id: 1, SkillName: "Анализ"},
		{
			Id: 2, SkillName: "Акробатика"}}, nil
}

type MockRepoError struct{}

func (m *MockRepoError) GetAllSkills(ctx context.Context) ([]Skills, error) {
	return nil, errors.New("mocked error")
}

type MockService struct{}

func (s *MockService) GetAllSkills(ctx context.Context) ([]Skills, error) {
	return []Skills{}, nil
}

type MockServiceError struct{}

func (s *MockServiceError) GetAllSkills(ctx context.Context) ([]Skills, error) {
	return nil, errors.New("fake service error")
}

func TestGetAllSkillsRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "skillName"}).
		AddRow(1, "Анализ").
		AddRow(2, "Акробатика")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, skillName FROM skills`)).WillReturnRows(mockRows)

	repo := NewRepository(db)

	ctx := context.Background()
	skills, err := repo.GetAllSkills(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(skills) != 2 {
		t.Errorf("expected 2 alignments, got %d", len(skills))
	}

	expected := []Skills{
		{Id: 1, SkillName: "Анализ"},
		{Id: 2, SkillName: "Акробатика"},
	}

	if !reflect.DeepEqual(skills, expected) {
		t.Errorf("expected %+v, got %+v", expected, skills)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllSkillsRepositoryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, skillName FROM skills`)).WillReturnError(sql.ErrConnDone)

	repo := NewRepository(db)

	ctx := context.Background()
	_, err = repo.GetAllSkills(ctx)
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
	skills, err := service.GetAllSkills(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expected := []Skills{
		{
			Id: 1, SkillName: "Анализ"},
		{
			Id: 2, SkillName: "Акробатика"}}
	if !reflect.DeepEqual(skills, expected) {
		t.Errorf("expected %+v, got %+v", expected, skills)
	}
}

func TestGetAllClassesErrorService(t *testing.T) {
	mockRepo := &MockRepoError{}
	service := NewService(mockRepo)

	ctx := context.Background()
	_, err := service.GetAllSkills(ctx)
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

	handler.GetAllSkills(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}

	var skills []Skills
	if err := json.Unmarshal(rr.Body.Bytes(), &skills); err != nil {
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

	handler.GetAllSkills(rr, req)

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

	handler.GetAllSkills(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expectedError := "fake service error"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}
