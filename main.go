package main

import (
	"estiam/library"
)

func main() {

	student1 := library.NewStudent("Léo")
	student2 := library.NewStudent("Pierre")

	professor1 := library.NewProfessor("Paul")
	professor2 := library.NewProfessor("Jack")

	users := make([]library.User, 0)
	users = append(users, student1)
	users = append(users, student2)
	users = append(users, professor1)
	users = append(users, professor2)

	library.PrintRent(&users)
}
