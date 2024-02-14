package alignment

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

func (m *MockRepo) GetAllAlignments(ctx context.Context) ([]Alignment, error) {
	return []Alignment{{Id: 1, AlignmentName: "Законопослушный"}, {Id: 2, AlignmentName: "Нейтральный"}}, nil
}

type MockRepoError struct{}

func (m *MockRepoError) GetAllAlignments(ctx context.Context) ([]Alignment, error) {
	return nil, errors.New("mocked error")
}

type MockService struct{}

func (s *MockService) GetAllAlignments(ctx context.Context) ([]Alignment, error) {
	return []Alignment{}, nil
}

type MockServiceError struct{}

func (s *MockServiceError) GetAllAlignments(ctx context.Context) ([]Alignment, error) {
	return nil, errors.New("fake service error")
}

func TestGetAllAlignmentsRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "alignmentName"}).
		AddRow(1, "Законопослушный").
		AddRow(2, "Нейтральный")

	mock.ExpectQuery("SELECT id, alignmentName FROM alignment").WillReturnRows(mockRows)

	repo := NewRepository(db)

	ctx := context.Background()
	alignments, err := repo.GetAllAlignments(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(alignments) != 2 {
		t.Errorf("expected 2 alignments, got %d", len(alignments))
	}

	expected := []Alignment{
		{Id: 1, AlignmentName: "Законопослушный"},
		{Id: 2, AlignmentName: "Нейтральный"},
	}

	for i, a := range alignments {
		if a.Id != expected[i].Id || a.AlignmentName != expected[i].AlignmentName {
			t.Errorf("expected %+v, got %+v", expected[i], a)
		}
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllAlignmentsRepository_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("SELECT id, alignmentName FROM alignment").WillReturnError(sql.ErrConnDone)

	repo := NewRepository(db)

	ctx := context.Background()
	_, err = repo.GetAllAlignments(ctx)
	if err == nil {
		t.Error("expected an error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAllAlignmentsService(t *testing.T) {
	mockRepo := &MockRepo{}
	service := NewService(mockRepo)

	ctx := context.Background()
	classes, err := service.GetAllAlignments(ctx)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expected := []Alignment{{Id: 1, AlignmentName: "Законопослушный"}, {Id: 2, AlignmentName: "Нейтральный"}}
	if !reflect.DeepEqual(classes, expected) {
		t.Errorf("expected %+v, got %+v", expected, classes)
	}
}

func TestGetAllAlignmentsErrorService(t *testing.T) {
	mockRepo := &MockRepoError{}
	service := NewService(mockRepo)

	ctx := context.Background()
	_, err := service.GetAllAlignments(ctx)
	if err == nil {
		t.Error("expected an error, got nil")
	} else if err.Error() != "mocked error" {
		t.Errorf("expected mocked error, got %v", err)
	}
}

func TestGetAllAlignmentsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllAlignments(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}

	var alignment []Alignment
	if err := json.Unmarshal(rr.Body.Bytes(), &alignment); err != nil {
		t.Errorf("error unmarshalling response body: %v", err)
	}
}

func TestGetAllAlignmentsHandler_Unauthorized(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 0, Err: errors.New("authorization header is missing")}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllAlignments(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}

	expectedError := "authorization header is missing"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}

func TestGetAllAlignmentsHandler_ServiceError(t *testing.T) {
	req, err := http.NewRequest("GET", "/classes", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer fake-token")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{}
	handler := NewHandler(svc, fakeTokenGetter)

	rr := httptest.NewRecorder()

	handler.GetAllAlignments(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	expectedError := "fake service error"
	if body := strings.TrimSpace(rr.Body.String()); body != expectedError {
		t.Errorf("handler returned unexpected body: got %v want %v", body, expectedError)
	}
}
