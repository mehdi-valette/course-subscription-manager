package routing

import (
	"app-htmx/repository"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

var students = []repository.Student{
	{
		Id:        1,
		Firstname: "Rachel",
		Lastname:  "Green",
		Email:     "rachel@green.net",
		Phone:     "000 343 43 34",
	},
	{
		Id:        2,
		Firstname: "Phoebe",
		Lastname:  "Buffay",
		Email:     "phoebe@buffay.net",
		Phone:     "342 423 49 80",
	},
}

func student(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/student/inline-form" {
		getStudentInlineForm(w, r)
		return
	}

	if r.URL.Path[0:17] == "/student/confirm/" {
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
	}
}

func getStudent(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("../templates/student-table.html")

	if err != nil {
		fmt.Println(err)
	}

	var studentList []repository.Student
	err = repository.GetStudentList(&studentList)

	if err != nil {
		fmt.Println(err)
	}

	tmp.Execute(w, studentList)
}

func getStudentInlineForm(w http.ResponseWriter, r *http.Request) {
	studentIdStr := r.URL.Query().Get("id")

	studentId64, err := strconv.ParseUint(studentIdStr, 10, 32)

	studentId := uint32(studentId64)
	var student repository.Student
	student.Id = studentId

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
	tmp, err := template.ParseFiles("../templates/student-table.html")

	if err != nil {
		fmt.Println(err)
	}

	id64, err := strconv.ParseUint(r.FormValue("Id"), 10, 32)

	if err != nil {
		fmt.Println(err)
	}

	id := uint32(id64)

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
	idRaw := strings.TrimSuffix(r.URL.Path[17:], "/")

	id64, err := strconv.ParseUint(idRaw, 10, 32)

	if err != nil {
		fmt.Println(err)
	}

	id := uint32(id64)
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

}
