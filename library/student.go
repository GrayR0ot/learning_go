package library

type Student struct {
	name string
}

func NewStudent(name string) *Student {
	return &Student{
		name,
	}
}
