package pages

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-9/jsonutil"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
)

func ChildrenPage(w http.ResponseWriter, r *http.Request) {
	IdStr := path.Base(r.URL.Path)
	ChildrenId, err := strconv.Atoi(IdStr)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	ClassInfoJSON, err := jsonutil.ReadJSONFile("class.json")
	if err != nil {
		log.Fatalf("Error reading a file: %s", err)
	}
	
	var class Class
	json.Unmarshal(ClassInfoJSON, &class)
	fmt.Fprintf(w, `<h1>Its a children page</h1><a href="/">Назад</a>`)
	for _, children := range class.Children {
		if children.Id == ChildrenId {
			fmt.Fprintf(w, `
		<ul>
			<li>Name: %s</li>
			<li>Age: %d</li>
			<li><a href='/children/%d'>ID: %d</a></li>
		</ul>
	`, children.Name, children.Age, children.Id, children.Id)
		}
	}
}
