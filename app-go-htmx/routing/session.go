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
	"Date":   func(date time.Time) string { return date.Format("Jan 02, 2006") },
	"String": func(date time.Time) string { return date.Format("2006-01-02") },
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

var sessionRowTemplate = template.Must(
	template.
		New("session-row").
		Funcs(templateDateFunction).
		ParseFiles("../templates/session-table.html"),
)

var sessionRowEditTemplate = template.Must(
	template.
		New("session-row-edit").
		Funcs(templateDateFunction).
		ParseFiles("../templates/session-table.html"),
)

func session(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/session/confirm/") {
		confirmDeleteSession(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "/session/session-row-edit/") {
		updateSessionForm(w, r)
		return
	}

	switch r.Method {
	case "GET":
		getSession(w, r)
	case "POST":
		addSession(w, r)
	case "PUT":
		updateSession(w, r)
	case "DELETE":
		deleteSession(w, r)
	}
}

func getSession(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/session/")

	if err == nil {
		session := repository.Session{ID: id}
		repository.GetSession(&session)
		sessionRowTemplate.Execute(w, session)
		return
	}

	var sessionList []repository.Session

	name := r.URL.Query().Get("name")
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := time.Parse("2006-01-02", startStr)

	if err != nil {
		start = time.Time{}
	}

	end, err := time.Parse("2006-01-02", endStr)

	if err != nil {
		end = time.Time{}
	}

	sessionFilter := repository.Session{Name: name, Start: start, End: end}

	repository.GetSessionList(&sessionList, sessionFilter)

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

func updateSessionForm(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/session/session-row-edit/")

	if handleError(w, err) {
		return
	}

	session := repository.Session{ID: id}

	if handleError(w, repository.GetSession(&session)) {
		return
	}

	sessionRowEditTemplate.Execute(w, session)
}

func updateSession(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/session/")

	if handleError(w, err) {
		return
	}

	name := r.FormValue("name")
	start, success := stringToTime(w, r.FormValue("start"))
	if !success {
		return
	}

	end, success := stringToTime(w, r.FormValue("end"))
	if !success {
		return
	}

	session := repository.Session{ID: id, Name: name, Start: start, End: end}

	repository.UpdateSession(&session)

	sessionRowNewTemplate.Execute(w, session)
}

/*
convert a string into a time

@param w: responseWriter to return the error

@param dateStr: date in the string format 2006-01-02

@return time: the converted time, or the zero time value

@return success: wheter the conversion was a success
*/
func stringToTime(w http.ResponseWriter, dateStr string) (time.Time, bool) {
	t, err := time.Parse("2006-01-02", dateStr)

	if handleError(w, err) {
		return time.Time{}, false
	}

	return t, true
}
