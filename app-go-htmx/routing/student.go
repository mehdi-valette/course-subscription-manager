package routing

import (
	"app-htmx/repository"
	"html/template"
	"io"
	"net/http"
	"strings"
)

const baseTemplate = "../templates/student-table.html"

var studentTableTemplate = template.Must(template.ParseFiles(baseTemplate))
var studentRowTemplate = template.Must(template.New("student-row").ParseFiles(baseTemplate))
var studentRowEditTemplate = template.Must(template.New("student-row-edit").ParseFiles(baseTemplate))
var studentRowNewTemplate = template.Must(template.New("student-row-new").ParseFiles(baseTemplate))
var studentDeleteConfirmTemplate = template.Must(template.New("student-delete-confirm").ParseFiles(baseTemplate))

func student(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/student/student-row-edit") {
		getStudentInlineForm(w, r)
		return
	}

	if len(r.URL.Path) > 17 && r.URL.Path[0:17] == "/student/confirm/" {
		getDeleteStudentConfirm(w, r)
		return
	}

	switch r.Method {
	case "GET":
		getStudent(w, r)
	case "PUT":
		updateStudent(w, r)
	case "POST":
		addStudent(w, r)
	case "DELETE":
		deleteStudent(w, r)
	}
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Path) > len("/student/") {
		getStudentSingle(w, r)
		return
	}

	getStudentList(w, r)
}

func getStudentList(w http.ResponseWriter, r *http.Request) {

	search := r.URL.Query().Get("search")

	var studentList []repository.Student

	var err error = nil
	if search != "" {
		err = repository.GetStudentFiltered(&studentList, search)
	} else {
		err = repository.GetStudentList(&studentList)
	}

	if handleError(w, err) {
		return
	}

	studentTableTemplate.Execute(w, studentList)
}

func getStudentSingle(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/student/")

	if handleError(w, err) {
		return
	}

	student := repository.Student{ID: id}
	err = repository.GetStudent(&student)

	if handleError(w, err) {
		return
	}

	studentRowTemplate.Execute(w, student)
}

func getStudentInlineForm(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/student/student-row-edit/")

	student := repository.Student{ID: id}

	err = repository.GetStudent(&student)

	if handleError(w, err) {
		return
	}

	studentRowEditTemplate.Execute(w, student)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/student/")

	if handleError(w, err) {
		return
	}

	student := repository.Student{
		ID:        id,
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone"),
	}

	err = repository.UpdateStudent(&student)

	if handleError(w, err) {
		return
	}

	studentRowTemplate.Execute(w, student)
}

func addStudent(w http.ResponseWriter, r *http.Request) {

	student := repository.Student{
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone"),
	}

	err := repository.AddStudent(&student)

	if handleError(w, err) {
		return
	}

	studentRowNewTemplate.Execute(w, student)
}

func getDeleteStudentConfirm(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/student/confirm/")

	if handleError(w, err) {
		return
	}

	student := repository.Student{ID: id}

	err = repository.GetStudent(&student)

	if handleError(w, err) {
		return
	}

	studentDeleteConfirmTemplate.Execute(w, student)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/student/")

	if handleError(w, err) {
		return
	}

	student := repository.Student{ID: id}

	err = repository.DeleteStudent(&student)

	if handleError(w, err) {
		return
	}

	io.WriteString(w, "")
}
