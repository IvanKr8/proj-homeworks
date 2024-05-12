package pages

import (
	"Users/user/Documents/Sites/projector/go-course/homeworks/hw-9/jsonutil"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	Id    int    `json:"Id"`
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Class string `json:"Class"`
}

type Class struct {
	Children []Student `json:"Children"`
}

func ClassInfo(w http.ResponseWriter, r *http.Request) {
	ClassInfoJSON, err := jsonutil.ReadJSONFile("class.json")
	if err != nil {
		log.Fatalf("Error reading a file: %s", err)
	}

	var class Class
	json.Unmarshal(ClassInfoJSON, &class)
	MainPage(w, r)
	for _, children := range class.Children {
		fmt.Fprintf(w, `
		<ul>
			<li>Name: %s</li>
			<li>Age: %d</li>
			<li><a href='/children/%d'>ID: %d</a></li>
		</ul>
	`, children.Name, children.Age, children.Id, children.Id)
	}
}
