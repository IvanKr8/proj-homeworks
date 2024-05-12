package login

import (
	"fmt"
	"net/http"
)

func Authorize(nextHandler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("teacher")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("checkToken(c.Value): ", checkToken(c.Value))
		if checkToken(c.Value) {
			nextHandler(w, r)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
