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

	if tmp, success := loadTemplateWith(w, "session-table.html", templateDateFunction); success {
		tmp.Execute(w, sessionList)
	}
}

func addSession(w http.ResponseWriter, r *http.Request) {

	start, startSuccess := stringToTime(w, r.FormValue("start"))
	end, endSuccess := stringToTime(w, r.FormValue("end"))

	if !startSuccess || !endSuccess {
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
