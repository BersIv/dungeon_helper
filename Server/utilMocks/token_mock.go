package utilMocks

import "net/http"

type MockTokenGetter struct {
	Id       int64
	Nickname string
	Err      error
}

func (t *MockTokenGetter) GetIdFromToken(r *http.Request) (int64, error) {
	return t.Id, t.Err
}

func (t *MockTokenGetter) GetNickNameFromToken(r *http.Request) (string, error) {
	return t.Nickname, t.Err
}
