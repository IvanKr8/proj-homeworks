package login

import (
	"fmt"
)

type Authorization struct {
	Login    string
	Password string
	Token    string
}

type Teacher struct {
	Authorization []Authorization
}

var teacher = Teacher{
	Authorization: []Authorization{
		{
			Login:    "user",
			Password: "password",
		},
	},
}

func checkToken(token string) bool {
	for _, val := range teacher.Authorization {
		fmt.Println(val.Token)
		fmt.Println("Checking token: ", token)
		if val.Token == token {
			return true
		}
	}

	return false
}

func authenticateUser(login, password string) (string, bool) {
	for i, val := range teacher.Authorization {
		if val.Login == login && val.Password == password {
			teacher.Authorization[i].Token = "TRUE"
			return teacher.Authorization[i].Token, true
		}
	}

	return "", false
}
