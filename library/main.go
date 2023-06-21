package main

import (
	"library/professor"
	"library/student"
	"library/user"
)

func main() {

	student1 := student.NewStudent("Léo")
	student2 := student.NewStudent("Pierre")

	professor1 := professor.NewProfessor("Paul")
	professor2 := professor.NewProfessor("Jack")

	users := make([]user.User, 0)
	users = append(users, student1)
	users = append(users, student2)
	users = append(users, professor1)
	users = append(users, professor2)

	user.PrintRent(&users)
}
