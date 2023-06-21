package student

import (
	"fmt"
)

type Student struct {
	Name string
}

func (s *Student) RentBook() {
	fmt.Printf("L'étudiant %s vient de réserver un livre.\n", s.Name)
}
