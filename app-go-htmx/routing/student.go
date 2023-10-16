package routing

import (
	"app-htmx/repository"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"
)

func student(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/student/inline-form") {
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
	tmp, err := template.ParseFiles("../templates/student-table.html")

	if err != nil {
		fmt.Println(err)
	}

	search := r.URL.Query().Get("search")

	var studentList []repository.Student

	if search != "" {
		err = repository.GetStudentFiltered(&studentList, search)
	} else {
		err = repository.GetStudentList(&studentList)
	}

	if err != nil {
		fmt.Println(err)
	}

	tmp.Execute(w, studentList)
}

func getStudentSingle(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/student/")

	if err != nil {
		tmp, errIntern := loadTemplate("index.html", w)
		if errIntern != nil {
			return
		}

		w.Header().Set("HX-Retarget", "#toast")
		w.Header().Set("HX-Reswap", "afterbegin")
		tmp.ExecuteTemplate(w, "error", ErrorTemplate{err.Error()})
		return
	}

	tmp, err := template.ParseFiles("../templates/student-table.html")

	if err != nil {
		fmt.Println(err)
	}

	student := repository.Student{Id: id}
	err = repository.GetStudent(&student)

	if err != nil {
		fmt.Println(err)
	}

	tmp.ExecuteTemplate(w, "inline-student", student)
}

func getStudentInlineForm(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/student/inline-form/")

	student := repository.Student{Id: id}

	err = repository.GetStudent(&student)

	if err != nil {
		fmt.Println(err)
	}

	tmp, err := template.ParseFiles("../templates/student-table.html")

	if err != nil {
		fmt.Println(err)
	}

	tmp.ExecuteTemplate(w, "inline-form", student)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/student/")

	if err != nil {
		fmt.Println(err)
	}

	tmp, err := template.ParseFiles("../templates/student-table.html")

	if err != nil {
		fmt.Println(err)
	}

	student := repository.Student{
		Id:        id,
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone"),
	}

	err = repository.UpdateStudent(&student)

	if err != nil {
		fmt.Println(err)
	}

	tmp.ExecuteTemplate(w, "inline-student", student)
}

func addStudent(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("../templates/student-table.html")

	if err != nil {
		fmt.Println(err)
	}

	student := repository.Student{
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
		Email:     r.FormValue("email"),
		Phone:     r.FormValue("phone"),
	}

	err = repository.AddStudent(&student)

	if err != nil {
		fmt.Println(err)
	}

	tmp.ExecuteTemplate(w, "inline-new-student", student)
}

func getDeleteStudentConfirm(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/student/confirm/")

	if err != nil {
		fmt.Println(err)
	}

	student := repository.Student{Id: id}

	err = repository.GetStudent(&student)

	if err != nil {
		fmt.Println(err)
	}

	tmp, err := template.ParseFiles("../templates/student-table.html")

	if err != nil {
		fmt.Println(err)
	}

	tmp.ExecuteTemplate(w, "student-delete-confirm", student)
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/student/")

	if err != nil {
		fmt.Println(err)
	}

	student := repository.Student{Id: id}

	err = repository.DeleteStudent(&student)

	if err != nil {
		fmt.Println(err)
	}

	io.WriteString(w, "")
}
