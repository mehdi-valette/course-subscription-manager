package routing

import (
	"errors"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type ErrorTemplate struct {
	Message string
}

var errorTemplate = template.Must(template.New("error").ParseFiles("../templates/index.html"))

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
write the error to the response when the error is not nil

@param w: ResponseWriter in which the error must be written

@param err: error to evaluate

@return errorHandled: true if there was an error, false otherwise
*/
func handleError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	w.Header().Set("HX-Retarget", "#toast")
	w.Header().Set("HX-Reswap", "afterbegin")
	errorTemplate.Execute(w, ErrorTemplate{err.Error()})

	return true
}
