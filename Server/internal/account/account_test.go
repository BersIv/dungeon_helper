package account

import (
	"bytes"
	"context"
	"dungeons_helper/utilMocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockRepo struct{}

func (m *MockRepo) CreateAccount(ctx context.Context, account *CreateAccountReq) error {
	return nil
}

func (m *MockRepo) GetAccountById(ctx context.Context, id int64) (*Account, error) {
	return &Account{}, nil
}

func (m *MockRepo) GetAccountByEmail(ctx context.Context, email string) (*Account, error) {
	return &Account{Id: 1, Email: "test@example.com", Nickname: "nickname", Avatar: "avatar_data"}, nil
}

func (m *MockRepo) UpdatePassword(ctx context.Context, account *Account) error {
	return nil
}

func (m *MockRepo) UpdateNickname(ctx context.Context, account *Account) error {
	return nil
}

func (m *MockRepo) UpdateAvatar(ctx context.Context, account *Account) error {
	return nil
}

type MockRepoError struct{}

func (m *MockRepoError) CreateAccount(ctx context.Context, account *CreateAccountReq) error {
	return errors.New("fake repository error")
}

func (m *MockRepoError) GetAccountById(ctx context.Context, id int64) (*Account, error) {
	return &Account{}, errors.New("fake repository error")
}

func (m *MockRepoError) GetAccountByEmail(ctx context.Context, email string) (*Account, error) {
	return &Account{}, errors.New("fake repository error")
}

func (m *MockRepoError) UpdatePassword(ctx context.Context, account *Account) error {
	return errors.New("fake repository error")
}

func (m *MockRepoError) UpdateNickname(ctx context.Context, account *Account) error {
	return errors.New("fake repository error")
}

func (m *MockRepoError) UpdateAvatar(ctx context.Context, account *Account) error {
	return errors.New("fake repository error")
}

type MockService struct{}

func (s *MockService) CreateAccount(c context.Context, req *CreateAccountReq) error {
	return nil
}

func (s *MockService) Login(c context.Context, req *LoginAccountReq) (*LoginAccountRes, error) {
	return &LoginAccountRes{Id: 1, Email: "test@example.com", Nickname: "nickname", Avatar: "avatar_data", accessToken: "000"}, nil
}

func (s *MockService) GoogleAuth(c context.Context, req *GoogleAcc) (*LoginAccountRes, error) {
	return &LoginAccountRes{Id: 1, Email: "test@example.com", Nickname: "nickname", Avatar: "avatar_data", accessToken: "000"}, nil
}

func (s *MockService) RestorePassword(c context.Context, req *RestoreReq) error {
	return nil
}

func (s *MockService) UpdateNickname(c context.Context, req *UpdateNicknameReq) error {
	return nil
}

func (s *MockService) UpdatePassword(c context.Context, req *UpdatePasswordReq) error {
	return nil
}

func (s *MockService) UpdateAvatar(c context.Context, req *UpdateAvatarReq) error {
	return nil
}

type MockServiceError struct{}

func (s *MockServiceError) CreateAccount(c context.Context, req *CreateAccountReq) error {
	return errors.New("fake service error")
}

func (s *MockServiceError) Login(c context.Context, req *LoginAccountReq) (*LoginAccountRes, error) {
	return &LoginAccountRes{}, errors.New("fake service error")
}

func (s *MockServiceError) GoogleAuth(c context.Context, req *GoogleAcc) (*LoginAccountRes, error) {
	return &LoginAccountRes{}, errors.New("fake service error")
}

func (s *MockServiceError) RestorePassword(c context.Context, req *RestoreReq) error {
	return errors.New("fake service error")
}

func (s *MockServiceError) UpdateNickname(c context.Context, req *UpdateNicknameReq) error {
	return errors.New("fake service error")
}

func (s *MockServiceError) UpdatePassword(c context.Context, req *UpdatePasswordReq) error {
	return errors.New("fake service error")
}

func (s *MockServiceError) UpdateAvatar(c context.Context, req *UpdateAvatarReq) error {
	return errors.New("fake service error")
}

func TestCreateAccountRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO image(image) VALUES (?)`)).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO account(email, password, nickname, idAvatar) VALUES (?, ?, ?, ?)`)).
		WithArgs("test@example.com", "password", "user", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateAccount(context.Background(), &CreateAccountReq{
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data",
	})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateAccountRepository_InsertImageError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO image(image) VALUES (?)`)).WillReturnError(errors.New("image insertion error"))

	err = repo.CreateAccount(context.Background(), &CreateAccountReq{
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data",
	})

	if err == nil {
		t.Error("expected error but got nil")
	} else {
		t.Logf("expected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateAccountRepository_InsertAccountError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO image(image) VALUES (?)`)).WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO account(email, password, nickname, idAvatar) VALUES (?, ?, ?, ?)`)).WillReturnError(errors.New("account insertion error"))

	err = repo.CreateAccount(context.Background(), &CreateAccountReq{
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data",
	})

	if err == nil {
		t.Error("expected error but got nil")
	} else {
		t.Logf("expected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCreateAccountService(t *testing.T) {
	mockRepo := &MockRepo{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.CreateAccount(ctx, &CreateAccountReq{
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data"})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestCreateAccountService_HashError(t *testing.T) {
	mockRepo := &MockRepo{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Err: errors.New("fake hash error")}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.CreateAccount(ctx, &CreateAccountReq{
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data"})
	if err.Error() != "fake hash error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestCreateAccountService_CreateError(t *testing.T) {
	mockRepo := &MockRepoError{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.CreateAccount(ctx, &CreateAccountReq{
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data"})
	if err == nil {
		t.Error("expected an error, got nil")
	} else if err.Error() != "fake repository error" {
		t.Errorf("expected fake repository error, got %v", err)
	}
}

func TestCreateAccountHandler(t *testing.T) {
	requestBody := []byte(`{
		"email": "1@mail.ru",
		"password": "1",
		"avatar": "123121",
		"nickname": "test"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}

	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.CreateAccount(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "application/json"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}
}

func TestCreateAccountHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{
		"email": "1@mail.ru",
		"password": "1",
		"avatar": "123121",
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.CreateAccount(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "text/plain; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}
}

func TestCreateAccountHandler_BadEmail(t *testing.T) {
	requestBody := []byte(`{
		"email": "1@",
		"password": "1",
		"avatar": "123121",
		"nickname": "test"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "123", Err: nil}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.CreateAccount(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "text/plain; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}
}

func TestCreateAccountHandler_ServiceError(t *testing.T) {
	requestBody := []byte(`{
		"email": "1@1.1",
		"password": "1",
		"avatar": "123121",
		"nickname": "test"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Authorization", "Bearer 1")

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}

	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.CreateAccount(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expectedContentType := "text/plain; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content-type header: got %v want %v", contentType, expectedContentType)
	}
}

func TestGetAccountByIdRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "email", "password", "nickname", "image"}).
		AddRow(1, "1@1.1", "123", "nickname", "image_data")

	mock.ExpectQuery(`SELECT a.id, email, password, nickname, i.image FROM account a 
						LEFT JOIN image i ON a.idAvatar = i.id 
						WHERE a.id = ?`).WillReturnRows(mockRows)

	repo := NewRepository(db)

	ctx := context.Background()
	account, err := repo.GetAccountById(ctx, 1)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if account == nil {
		t.Errorf("expected account data, got %+v", account)
	}

	expected := &Account{
		Id: 1, Email: "1@1.1", Password: "123", Nickname: "nickname", Avatar: "image_data",
	}

	if !reflect.DeepEqual(account, expected) {
		t.Errorf("expected %+v, got %+v", expected, account)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAccountByIdRepository_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT a.id, email, password, nickname, i.image FROM account a 
						LEFT JOIN image i ON a.idAvatar = i.id 
						WHERE a.id = ?`).WillReturnError(errors.New("No rows"))

	repo := NewRepository(db)

	ctx := context.Background()
	_, err = repo.GetAccountById(ctx, 1)
	if err == nil {
		t.Error("expected an error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAccountByEmailRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockRows := sqlmock.NewRows([]string{"id", "email", "password", "nickname", "image"}).
		AddRow(1, "1@1.1", "123", "nickname", "image_data")

	mock.ExpectQuery(`SELECT a.id, email, password, nickname, i.image FROM account a 
						LEFT JOIN image i ON a.idAvatar = i.id 
						WHERE email = ?`).WillReturnRows(mockRows)

	repo := NewRepository(db)

	ctx := context.Background()
	account, err := repo.GetAccountByEmail(ctx, "1@1.1")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if account == nil {
		t.Errorf("expected account data, got %+v", account)
	}

	expected := &Account{
		Id: 1, Email: "1@1.1", Password: "123", Nickname: "nickname", Avatar: "image_data",
	}

	if !reflect.DeepEqual(account, expected) {
		t.Errorf("expected %+v, got %+v", expected, account)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetAccountByEmailRepository_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery(`SELECT a.id, email, password, nickname, i.image FROM account a 
						LEFT JOIN image i ON a.idAvatar = i.id 
						WHERE email = ?`).WillReturnError(errors.New("No rows"))

	repo := NewRepository(db)

	ctx := context.Background()
	_, err = repo.GetAccountByEmail(ctx, "1@1.1")
	if err == nil {
		t.Error("expected an error, got nil")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestLoginService(t *testing.T) {
	mockRepo := &MockRepo{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	classes, err := service.Login(ctx, &LoginAccountReq{
		Email:    "test@example.com",
		Password: "password"})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expected := &LoginAccountRes{Id: 1, Email: "test@example.com", Nickname: "nickname", Avatar: "avatar_data", accessToken: "000"}
	if classes.accessToken == "" {
		t.Errorf("expected non-empty accessToken, got empty")
	}
	if !reflect.DeepEqual(classes.Id, expected.Id) ||
		!reflect.DeepEqual(classes.Email, expected.Email) ||
		!reflect.DeepEqual(classes.Nickname, expected.Nickname) ||
		!reflect.DeepEqual(classes.Avatar, expected.Avatar) {
		t.Errorf("expected %+v, got %+v", expected, classes)
	}
}

func TestLoginService_ErrorRepo(t *testing.T) {
	mockRepo := &MockRepoError{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	_, err := service.Login(ctx, &LoginAccountReq{
		Email:    "test@example.com",
		Password: "password"})
	if err.Error() != "fake repository error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestLoginService_PasswordError(t *testing.T) {
	mockRepo := &MockRepo{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "", Err: errors.New("bad password error")}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	_, err := service.Login(ctx, &LoginAccountReq{
		Email:    "test@example.com",
		Password: "password"})
	if err.Error() != "bad password error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestLoginHandler(t *testing.T) {
	requestBody := []byte(`{
		"email": "1@mail.ru",
		"password": "1"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.Login(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	cookie := rr.Result().Cookies()[0]
	if cookie.Name != "jwt" || cookie.Value != "000" {
		t.Errorf("unexpected cookie: %s=%s", cookie.Name, cookie.Value)
	}
}

func TestLoginHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{
		"email": "",
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.Login(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestLoginHandler_StatusUnauthorized(t *testing.T) {
	requestBody := []byte(`{
		"email": "1@mail.ru",
		"password": "1"
		}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.Login(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestLogoutHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h := &Handler{}
		h.Logout(w, r)
	})

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("expected Content-Type %s, got %s", "application/json", contentType)
	}

	cookies := recorder.Result().Cookies()
	if len(cookies) != 1 {
		t.Errorf("expected 1 cookie, got %d", len(cookies))
	}
	cookie := cookies[0]
	if cookie.Name != "jwt" || cookie.Value != "" || cookie.MaxAge != -1 {
		t.Errorf("unexpected cookie: %v", cookie)
	}

	expectedBody := `{"message": "logout successful"}`
	if recorder.Body.String() != expectedBody {
		t.Errorf("expected body %s, got %s", expectedBody, recorder.Body.String())
	}
}

func TestGoogleAuthService(t *testing.T) {
	mockRepo := &MockRepo{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	classes, err := service.GoogleAuth(ctx, &GoogleAcc{
		Email:   "test@example.com",
		Picture: "picture_data"})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	expected := &LoginAccountRes{Id: 1, Email: "test@example.com", Nickname: "nickname", Avatar: "avatar_data", accessToken: "000"}
	if classes.accessToken == "" {
		t.Errorf("expected non-empty accessToken, got empty")
	}
	if !reflect.DeepEqual(classes.Id, expected.Id) ||
		!reflect.DeepEqual(classes.Email, expected.Email) ||
		!reflect.DeepEqual(classes.Nickname, expected.Nickname) ||
		!reflect.DeepEqual(classes.Avatar, expected.Avatar) {
		t.Errorf("expected %+v, got %+v", expected, classes)
	}
}

func TestGoogleAuthService_RepoError(t *testing.T) {
	mockRepo := &MockRepoError{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	_, err := service.GoogleAuth(ctx, &GoogleAcc{
		Email:   "test@example.com",
		Picture: "picture_data"})
	if err.Error() != "fake repository error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestLoginGoogleHandler(t *testing.T) {
	requestBody := []byte(`{"token": "token_data"}`)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}

	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	handler.LoginGoogle(recorder, req)

	if recorder.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, recorder.Code)
	}
}

func TestLoginGoogleHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{"token": `)
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}

	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	handler.LoginGoogle(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}
}

func TestUpdatePasswordRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE account SET password = ? WHERE id = ?`)).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdatePassword(context.Background(), &Account{
		Id:       1,
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data",
	})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdatePasswordRepository_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE account SET password = ? WHERE id = ?`)).WillReturnError(errors.New("update error"))

	err = repo.UpdatePassword(context.Background(), &Account{
		Id:       1,
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data",
	})
	if err.Error() != "update error" {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdatePasswordService(t *testing.T) {
	mockRepo := &MockRepo{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.UpdatePassword(ctx, &UpdatePasswordReq{
		Id:          1,
		NewPassword: "password"})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestUpdatePasswordService_HashError(t *testing.T) {
	mockRepo := &MockRepo{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: errors.New("fake hash error")}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.UpdatePassword(ctx, &UpdatePasswordReq{
		Id:          1,
		NewPassword: "password"})
	if err.Error() != "fake hash error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestUpdatePasswordService_ErrorRepo(t *testing.T) {
	mockRepo := &MockRepoError{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.UpdatePassword(ctx, &UpdatePasswordReq{
		Id:          1,
		NewPassword: "password"})
	if err.Error() != "fake repository error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestRestorePasswordService(t *testing.T) {
	mockRepo := &MockRepo{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.RestorePassword(ctx, &RestoreReq{
		Email:       "1@1.1",
		NewPassword: "password"})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestRestorePasswordService_ErrorRepo(t *testing.T) {
	mockRepo := &MockRepoError{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.RestorePassword(ctx, &RestoreReq{
		Email:       "1@1.1",
		NewPassword: "password"})
	if err.Error() != "fake repository error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestUpdatePasswordHandler(t *testing.T) {
	requestBody := []byte(`{
		"id": 1,
		"newPassword": "newPassword"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdatePassword(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestUpdatePasswordHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{
		"": dsds
		"newPassword": "newPassword"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdatePassword(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestUpdatePasswordHandler_StatusUnauthorized(t *testing.T) {
	requestBody := []byte(`{
		"id": 1,
		"newPassword": "newPassword"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: errors.New("no token")}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdatePassword(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, rr.Code)
	}
}

func TestUpdatePasswordHandler_ErrorService(t *testing.T) {
	requestBody := []byte(`{
		"id": 1,
		"newPassword": "newPassword"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdatePassword(rr, req)
	body := rr.Body.String()
	if body != "fake service error\n" {
		t.Errorf("expected fake service error, got %+v", body)
	}
}

func TestRestorePasswordHandler(t *testing.T) {
	requestBody := []byte(`{
		"email": "1@1.1",
		"newPassword": "newPassword"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.RestorePassword(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestRestorePasswordHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{
		"email": "1@1,
		"newPassword": "newPassword"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.RestorePassword(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestRestorePasswordHandler_EmptyEmail(t *testing.T) {
	requestBody := []byte(`{
		"email": "",
		"newPassword": "newPassword"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.RestorePassword(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}

func TestRestorePasswordHandler_ErrorService(t *testing.T) {
	requestBody := []byte(`{
		"email": "1@1.1",
		"newPassword": "newPassword"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.RestorePassword(rr, req)

	body := rr.Body.String()
	if body != "fake service error\n" {
		t.Errorf("expected fake service error, got %+v", body)
	}
}

func TestSendEmail_AuthenticationError(t *testing.T) {
	os.Setenv("SMTP_HOST", "test.smtp.com")
	os.Setenv("SMTP_PORT", "587")
	os.Setenv("SMTP_USERNAME", "testuser")
	os.Setenv("SMTP_PASSWORD", "invalidpassword")

	toEmail := "1@1.1"
	subject := "Test Subject"
	body := "Test Body"

	err := SendEmail(toEmail, subject, body)
	if err == nil {
		t.Errorf("expected authentication error, got nil")
	}
}

func TestWelcomeEmail_AuthenticationError(t *testing.T) {
	os.Setenv("SMTP_HOST", "test.smtp.com")
	os.Setenv("SMTP_PORT", "587")
	os.Setenv("SMTP_USERNAME", "testuser")
	os.Setenv("SMTP_PASSWORD", "invalidpassword")

	toEmail := "1@1.1"

	err := SendWelcomeEmail(toEmail)
	if err == nil {
		t.Errorf("expected authentication error, got nil")
	}
}

func TestUpdateNicknameRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE account SET nickname = ? WHERE id = ?`)).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateNickname(context.Background(), &Account{
		Id:       1,
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data",
	})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateNicknameRepository_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE account SET nickname = ? WHERE id = ?`)).WillReturnError(errors.New("update error"))

	err = repo.UpdateNickname(context.Background(), &Account{
		Id:       1,
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data",
	})
	if err.Error() != "update error" {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateNicknameService(t *testing.T) {
	mockRepo := &MockRepo{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.UpdateNickname(ctx, &UpdateNicknameReq{
		Id:       1,
		Nickname: "newNickname"})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestUpdateNicknameService_ErrorRepo(t *testing.T) {
	mockRepo := &MockRepoError{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.UpdateNickname(ctx, &UpdateNicknameReq{
		Id:       1,
		Nickname: "newNickname"})
	if err.Error() != "fake repository error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestUpdateNicknameHandler(t *testing.T) {
	requestBody := []byte(`{
		"id": 1,
		"nickname": "newNickname"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdateNickname(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestUpdateNicknameHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{
		"": dsds
		"nickname": "newNickname"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdateNickname(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestUpdateNicknameHandler_StatusUnauthorized(t *testing.T) {
	requestBody := []byte(`{
		"id": 1,
		"newPassword": "newPassword"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: errors.New("no token")}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdateNickname(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, rr.Code)
	}
}

func TestUpdateNicknameHandler_ErrorService(t *testing.T) {
	requestBody := []byte(`{
		"id": 1,
		"newPassword": "newPassword"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdateNickname(rr, req)
	body := rr.Body.String()
	if body != "fake service error\n" {
		t.Errorf("expected fake service error, got %+v", body)
	}
}

func TestUpdateAvatarRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE image SET image = ? WHERE id = (SELECT idAvatar FROM account WHERE id = ?)`)).WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateAvatar(context.Background(), &Account{
		Id:       1,
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data",
	})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateAvatarRepository_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewRepository(db)

	mock.ExpectExec(regexp.QuoteMeta(`UPDATE image SET image = ? WHERE id = (SELECT idAvatar FROM account WHERE id = ?)`)).WillReturnError(errors.New("update error"))

	err = repo.UpdateAvatar(context.Background(), &Account{
		Id:       1,
		Email:    "test@example.com",
		Password: "password",
		Nickname: "user",
		Avatar:   "avatar_data",
	})
	if err.Error() != "update error" {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdatAvatarService(t *testing.T) {
	mockRepo := &MockRepo{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.UpdateAvatar(ctx, &UpdateAvatarReq{
		Id:     1,
		Avatar: "avatar_data"})
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestUpdateAvatarService_ErrorRepo(t *testing.T) {
	mockRepo := &MockRepoError{}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{Password: "password", Err: nil}
	service := NewService(mockRepo, mockedPasswordHasher)

	ctx := context.Background()
	err := service.UpdateAvatar(ctx, &UpdateAvatarReq{
		Id:     1,
		Avatar: "avatar_data"})
	if err.Error() != "fake repository error" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestUpdateAvatarHandler(t *testing.T) {
	requestBody := []byte(`{
		"id": 1,
		"avatar": "avatar_data"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdateAvatar(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestUpdateAvatarHandler_BadRequest(t *testing.T) {
	requestBody := []byte(`{
		"": dsds
		"avatar": "avatar_data"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdateAvatar(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
}

func TestUpdateAvatarHandler_StatusUnauthorized(t *testing.T) {
	requestBody := []byte(`{
		"id": 1,
		"avatar": "avatar_data"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockService{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: errors.New("no token")}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdateAvatar(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected status %d, got %d", http.StatusUnauthorized, rr.Code)
	}
}

func TestUpdateAvatarHandler_ErrorService(t *testing.T) {
	requestBody := []byte(`{
		"id": 1,
		"avatar": "avatar_data"
	}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatal(err)

	}

	svc := &MockServiceError{}
	fakeTokenGetter := &utilMocks.MockTokenGetter{Id: 123, Err: nil}
	mockedPasswordHasher := &utilMocks.MockPasswordHasher{}
	handler := NewHandler(svc, fakeTokenGetter, mockedPasswordHasher)

	rr := httptest.NewRecorder()
	handler.UpdateAvatar(rr, req)
	body := rr.Body.String()
	if body != "fake service error\n" {
		t.Errorf("expected fake service error, got %+v", body)
	}
}
