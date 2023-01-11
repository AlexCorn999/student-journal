package students_repo

import (
	"fmt"
	"students/pkg/student"
)

// интерфейс представление класса школы у которого есть медоты:
// вывести всех студентов, добавить нового студента и получить информацию о студенте.
type AllClasses interface {
	StudentsInfo()
	Get(name string) (*student.Student, error)
	Put(s *student.Student)
}

type StudentsStorage map[string]*student.Student

// создание хранилища
func NewStudentsStorage() StudentsStorage {
	return make(map[string]*student.Student)
}

// добавляет студента в карту
func (ss StudentsStorage) Put(s *student.Student) {
	ss[s.Name] = s
}

// выводит информацию одного студента
func (ss StudentsStorage) Get(name string) (*student.Student, error) {
	s, ok := ss[name]
	if !ok {
		return nil, fmt.Errorf("no such student")
	}
	return s, nil
}

// выводит информацию всех студентов
func (ss StudentsStorage) StudentsInfo() {
	fmt.Println("\nСтуденты из хранилища:")
	for _, s := range ss {
		fmt.Println(*s)
	}
}
