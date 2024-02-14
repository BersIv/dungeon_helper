package utilMocks

type MockPasswordHasher struct {
	Password string
	Err      error
}

func (ph *MockPasswordHasher) HashPassword(password string) (string, error) {
	return ph.Password, ph.Err
}

func (ph *MockPasswordHasher) CheckPassword(password string, hashedPassword string) error {
	return ph.Err
}

func (ph *MockPasswordHasher) GeneratePassword() string {
	return ph.Password
}
