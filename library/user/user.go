package user

type User interface {
	RentBook( /*b *book*/ )
}

func PrintRent(users *[]User) {
	for _, u := range *users {
		u.RentBook()
	}
}
