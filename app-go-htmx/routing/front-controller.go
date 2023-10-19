package routing

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type User struct {
	User string
	List []struct {
		Name string
		Done bool
	}
}

func FrontController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		getIndex(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/student") {
		student(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/session") {
		session(w, r)
	}
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("../templates/index.html")

	if err != nil {
		fmt.Println(err)
	}

	err = tmp.Execute(w, nil)

	if err != nil {
		fmt.Println(err)
	}
}
