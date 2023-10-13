package repository

func AddStudent(student *Student) error {
	result := db.Create(student)

	return result.Error
}

func GetStudentList(studentList *[]Student) error {
	result := db.Order("firstname, lastname").Find(studentList)

	return result.Error
}

func GetStudentFiltered(studentList *[]Student, filter string) error {
	result := db.Order("firstname, lastname").Where("firstname || lastname || email || phone LIKE ?", "%"+filter+"%").Find(studentList)

	return result.Error
}

func GetStudent(student *Student) error {
	result := db.First(student, student.Id)

	return result.Error
}

func UpdateStudent(student *Student) error {
	result := db.Save(student)

	return result.Error
}

func DeleteStudent(student *Student) error {
	result := db.Delete(student)

	return result.Error
}
