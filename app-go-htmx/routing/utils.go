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

/*
get the ID at the end of a path

@param path: the whole path requested by the client

@param commonPath: the part of the path before the ID

@return id: the ID that was in the path

@return error: the error that occurred, or nil when no error occurred
*/
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

/*
load a template

@param w: ResponseWriter to write the error to

@param path: path of the template, starting from the directory "template"

@return template: the parsed template

@return success: true if the parsing was successful, false otherwise
*/
func loadTemplate(w http.ResponseWriter, path string) (*template.Template, bool) {
	tmp, err := template.ParseFiles("../templates/" + path)

	if handleError(w, err) {
		return nil, false
	}

	return tmp, true
}

/*
write the error to the response when the error is not nil

@param w: ResponseWriter in which the error must be written

@param err: error to evaluate

@return errorHandled: true if there was an error, false otherwise
*/
func handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	if tmp, success := loadTemplate(w, "index.html"); success {
		w.Header().Set("HX-Retarget", "#toast")
		w.Header().Set("HX-Reswap", "afterbegin")
		tmp.ExecuteTemplate(w, "error", ErrorTemplate{err.Error()})

		return true
	}

	w.Header().Set("HX-Retarget", "#toast")
	w.Header().Set("HX-Reswap", "afterbegin")
	io.WriteString(w, "<div class=\"toast\" style=\"background-color: red;\">"+err.Error()+"</div>")

	return true
}
