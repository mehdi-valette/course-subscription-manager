package routing

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type ErrorTemplate struct {
	Message string
}

func extractIdFromPath(path string, commonPath string) (uint32, error) {
	if len(path) < len(commonPath) {
		return 0, errors.New("the path is invalid")
	}

	idRaw := strings.TrimSuffix(path[len(commonPath):], "/")

	id64, err := strconv.ParseUint(idRaw, 10, 32)

	if err != nil {
		return 0, err
	}

	return uint32(id64), nil
}

func loadTemplate(path string, w http.ResponseWriter) (*template.Template, error) {
	tmp, err := template.ParseFiles("../templates/" + path)

	if err != nil {
		io.WriteString(w, err.Error())
		return nil, err
	}

	return tmp, nil
}

func handleError(w http.ResponseWriter, err error) bool {
	if err != nil {
		tmp, errIntern := loadTemplate("index.html", w)
		if errIntern != nil {
			w.Header().Set("HX-Retarget", "#toast")
			w.Header().Set("HX-Reswap", "afterbegin")
			io.WriteString(w, "<div class=\"toast\">"+err.Error()+"</div>")
			return true
		}

		w.Header().Set("HX-Retarget", "#toast")
		w.Header().Set("HX-Reswap", "afterbegin")
		tmp.ExecuteTemplate(w, "error", ErrorTemplate{err.Error()})
		return true
	}

	return false
}
