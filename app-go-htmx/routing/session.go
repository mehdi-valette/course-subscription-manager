package routing

import (
	"app-htmx/repository"
	"html/template"
	"net/http"
	"time"
)

var templateDateFunction = template.FuncMap{
	"Date": func(date time.Time) string { return date.Format("Jan 02, 2006") },
}

var sessionTableTemplate = template.Must(
	template.
		New("session-table.html").
		Funcs(
			template.FuncMap{"Date": func(date time.Time) string {
				return date.Format("Jan 02, 2006")
			},
			},
		).
		ParseFiles("../templates/session-table.html"))

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

	err := repository.AddSession(&session)
	if handleError(w, err) {
		return
	}
}

func stringToTime(w http.ResponseWriter, dateStr string) (time.Time, bool) {
	t, err := time.Parse("2006-01-02", dateStr)

	if handleError(w, err) {
		return time.Time{}, false
	}

	return t, true
}
