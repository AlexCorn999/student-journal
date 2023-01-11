package student

type Student struct {
	Name  string
	Age   int
	Grade int
}

// создает структуру со студентом
func NewStudent(name string, age int, grade int) *Student {
	return &Student{name, age, grade}
}
