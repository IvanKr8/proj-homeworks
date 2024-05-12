package pages

import (
	"fmt"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>Its a main page</h1>
	<a href='/'>main</a><br>
	<a href='/class-info'>Class Info</a><br>
	<a href='/login'>I'm teacher!'</a><br>
	`)
}
