package routing

import (
	"app-htmx/repository"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		getSession(w, r)
	case "POST":
		addSession(w, r)
	}
}

func getSession(w http.ResponseWriter, r *http.Request) {
	var sessionList []repository.Session
	repository.GetSessionList(&sessionList)

	if tmp, success := loadTemplate(w, "session-table.html"); success {
		tmp.Execute(w, sessionList)
	}
}

func addSession(w http.ResponseWriter, r *http.Request) {
	session := repository.Session{
		Name: r.FormValue("name"),
	}

	err := repository.AddSession(&session)
	if handleError(w, err) {
		return
	}
}
