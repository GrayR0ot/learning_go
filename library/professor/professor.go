package professor

type Professor struct {
	name string
}

func NewProfessor(name string) *Professor {
	return &Professor{
		name,
	}
}
