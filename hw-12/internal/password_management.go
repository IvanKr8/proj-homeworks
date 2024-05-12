package internal

func (pm *PasswordManager) AddPassword(name, password string) {
	pm.password[name] = password
}

func (pm *PasswordManager) ListPassword() []string {
	names := make([]string, len(pm.password))
	for name := range pm.password {
		names = append(names, name)
	}
	return names
}

func (pm *PasswordManager) GetPassword(name string) (string, bool) {
	password, ok := pm.password[name]
	return password, ok
}
