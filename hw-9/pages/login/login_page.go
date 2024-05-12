package login

import (
	"fmt"
	"net/http"
	"time"
)

func LoginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		login := r.FormValue("login")
		password := r.FormValue("password")
		token, ok := authenticateUser(login, password)
		if ok {
			http.SetCookie(w, &http.Cookie{
				Name:    "teacher",
				Value:   token,
				Expires: time.Now().Add(time.Minute),
			})
			fmt.Println("Cookie set: ", token)

			http.Redirect(w, r, "/", http.StatusFound)

		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}

	fmt.Fprintf(w, `<h1>Login</h1>
    <form action="/login" method="post">
        <label for="login">Логин:</label><br>
        <input type="text" id="login" name="login"><br>
        <label for="password">Пароль:</label><br>
        <input type="password" id="password" name="password"><br>
        <input type="submit" value="Отправить">
    </form>`)

}
