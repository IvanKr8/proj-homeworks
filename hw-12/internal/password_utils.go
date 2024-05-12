package internal

type PasswordManager struct {
	password map[string]string
}

func NewPasswordManager() *PasswordManager {
	return &PasswordManager{password: make(map[string]string)}
}
