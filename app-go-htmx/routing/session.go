package routing

import (
	"app-htmx/repository"
	"html/template"
	"io"
	"net/http"
	"strings"
	"time"
)

var templateDateFunction = template.FuncMap{
	"Date": func(date time.Time) string { return date.Format("Jan 02, 2006") },
}

var sessionTableTemplate = template.Must(
	template.
		New("session-table.html").
		Funcs(templateDateFunction).
		ParseFiles("../templates/session-table.html"),
)

var sessionConfirmTemplate = template.Must(
	template.
		New("session-delete-confirm").
		Funcs(templateDateFunction).
		ParseFiles("../templates/session-table.html"),
)

var sessionRowNewTemplate = template.Must(
	template.
		New("session-row-new").
		Funcs(templateDateFunction).
		ParseFiles("../templates/session-table.html"),
)

func session(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/session/confirm/") {
		confirmDeleteSession(w, r)
		return
	}

	switch r.Method {
	case "GET":
		getSession(w, r)
	case "POST":
		addSession(w, r)
	case "DELETE":
		deleteSession(w, r)
	}
}

func getSession(w http.ResponseWriter, r *http.Request) {
	var sessionList []repository.Session
	repository.GetSessionList(&sessionList)

	sessionTableTemplate.Execute(w, sessionList)
}

func addSession(w http.ResponseWriter, r *http.Request) {

	start, success := stringToTime(w, r.FormValue("start"))

	if !success {
		return
	}

	end, success := stringToTime(w, r.FormValue("end"))

	if !success {
		return
	}

	session := repository.Session{
		Name:  r.FormValue("name"),
		Start: start,
		End:   end,
	}

	if handleError(w, repository.AddSession(&session)) {
		return
	}

	sessionRowNewTemplate.Execute(w, session)
}

func confirmDeleteSession(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/session/confirm/")

	if handleError(w, err) {
		return
	}

	session := repository.Session{ID: id}

	if handleError(w, repository.GetSession(&session)) {
		return
	}

	sessionConfirmTemplate.Execute(w, session)
}

func deleteSession(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/session/")

	if handleError(w, err) {
		return
	}

	session := repository.Session{ID: id}

	if handleError(w, repository.DeleteSession(&session)) {
		return
	}

	io.WriteString(w, "")
}

func stringToTime(w http.ResponseWriter, dateStr string) (time.Time, bool) {
	t, err := time.Parse("2006-01-02", dateStr)

	if handleError(w, err) {
		return time.Time{}, false
	}

	return t, true
}
