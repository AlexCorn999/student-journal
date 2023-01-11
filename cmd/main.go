package main

import (
	"fmt"
	"io"
	"students/pkg/student"
	students_repo "students/pkg/student_repo"
)

func main() {
	var name string
	var age int
	var grade int

	nineClass := students_repo.NewStudentsStorage()
	// к выпускному классу присваиваем 9 класс
	var graduationClass students_repo.AllClasses = nineClass

	for {
		fmt.Print("Введите данные нового студента (Имя,Возраст,Оценку) :")
		_, err := fmt.Scan(&name, &age, &grade)
		if err == io.EOF {
			break
		}
		nineClass.Put(student.NewStudent(name, age, grade))
	}

	// получаем информацию о выпускном классе
	graduationClass.StudentsInfo()

	// добавили нового студента в выпускной класс
	graduationClass.Put(student.NewStudent("Kateeeeee", 44, 2))
	graduationClass.StudentsInfo()

	getName, err := graduationClass.Get("Zik")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(*getName)
	}

}
