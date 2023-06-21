package main

import (
	"library/professor"
	"library/student"
	"library/user"
)

func main() {

	student1 := student.Student{Name: "Léo"}
	student2 := student.Student{Name: "Pierre"}

	professor1 := professor.Professor{Name: "Paul"}
	professor2 := professor.Professor{Name: "Jack"}

	users := make([]user.User, 0)
	users = append(users, student1)
	users = append(users, student2)
	users = append(users, professor1)
	users = append(users, professor2)

	user.PrintRent(&users)
}
