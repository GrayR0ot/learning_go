package professor

import "fmt"

type Professor struct {
	Name string
}

func (p *Professor) RentBook() {
	fmt.Printf("L'enseignant %s vient de réserver un livre.\n", p.Name)
}
