package util

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJWTTokenGetterGetIdFromTokenExpired(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywiaXNzIjoiMyIsImV4cCI6MTcwNTg1NTU0NX0.uLh7ejjLeZ4f0iuxNrGvgorkn3JJv2zKjZNfrQTAsS0")

	tokenGetter := JWTTokenGetter{}

	_, err := tokenGetter.GetIdFromToken(req)

	if err.Error() != "token signature is invalid: signature is invalid" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestJWTTokenGetter_GetIdFromToken_AuthorizationHeaderMissing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	tokenGetter := JWTTokenGetter{}

	_, err := tokenGetter.GetIdFromToken(req)

	if err.Error() != "authorization header is missing" {
		t.Errorf("expected authorization header is missing, got %v", err)
	}
}

func TestJWTTokenGetterGetIdFromTokenInvalidTokenFormat(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "invalid_token_format")

	tokenGetter := JWTTokenGetter{}

	_, err := tokenGetter.GetIdFromToken(req)

	if err.Error() != "invalid token format" {
		t.Errorf("expected invalid token format, got %v", err)
	}
}

func TestJWTTokenGetter_GetNickNameFromToken_Success(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MywiaXNzIjoiMyIsImV4cCI6MTcwNTg1NTU0NX0.uLh7ejjLeZ4f0iuxNrGvgorkn3JJv2zKjZNfrQTAsS0")

	tokenGetter := JWTTokenGetter{}

	_, err := tokenGetter.GetNickNameFromToken(req)

	if err.Error() != "token signature is invalid: signature is invalid" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestJWTTokenGetter_GetNickNameFromToken_AuthorizationHeaderMissing(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	tokenGetter := JWTTokenGetter{}

	_, err := tokenGetter.GetNickNameFromToken(req)

	expectedError := "authorization header is missing"
	if err.Error() != expectedError {
		t.Errorf("expected error: %s, got: %v", expectedError, err)
	}
}

func TestJWTTokenGetter_GetNickNameFromToken_InvalidTokenFormat(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "invalid_token_format")

	tokenGetter := JWTTokenGetter{}

	_, err := tokenGetter.GetNickNameFromToken(req)

	expectedError := "invalid token format"
	if err.Error() != expectedError {
		t.Errorf("expected error: %s, got: %v", expectedError, err)
	}
}

func TestPasswordHasher(t *testing.T) {
	passwordHasher := BcryptPasswordHasher{}

	password := passwordHasher.GeneratePassword()

	hashedPassword, err := passwordHasher.HashPassword(password)
	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	err = passwordHasher.CheckPassword(password, hashedPassword)
	if err != nil {
		t.Fatalf("password and hashed password do not match: %v", err)
	}

	wrongPassword := "wrong_password"
	err = passwordHasher.CheckPassword(wrongPassword, hashedPassword)
	if err == nil {
		t.Fatalf("expected password and hashed password do not match error, but got nil")
	}
}

func TestLoggingMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	var buf bytes.Buffer

	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	rr := httptest.NewRecorder()

	loggingMiddleware := LoggingMiddleware(handler)
	loggingMiddleware.ServeHTTP(rr, req)

	if buf.String() != "" {
		t.Errorf("unexpected output: got %q", buf.String())
	}
}
