package user

import (
	"fmt"
)

type User interface {
	RentBook( /*b *book*/ )
}

func (p *professor.Professor) RentBook() {
	fmt.Printf("L'enseignant %s vient de réserver un livre.\n", p.name)
}

func (s *student.Student) RentBook() {
	fmt.Printf("L'étudiant %s vient de réserver un livre.\n", s.name)
}

func PrintRent(users *[]User) {
	for _, u := range *users {
		u.RentBook()
	}
}
