//METHODS

// Programme d'exemple montrant comment déclarer des méthodes et comment le compilateur Go les prend en charge.
package main

import (
	"fmt"
)

// user définit un utilisateur dans le programme.
type user struct {
	name  string
	email string
}

// notify implémente une méthode avec un récepteur de valeur.
func (u user) notify() {
	fmt.Printf("Utilisateur : Envoi de l'e-mail de l'utilisateur à %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implémente une méthode avec un récepteur de pointeur.
func (u *user) changeEmail(email string) {
	u.email = email
}

// main est le point d'entrée de l'application.
func main() {
	// Les valeurs de type user peuvent être utilisées pour appeler des méthodes
	// déclarées avec un récepteur de valeur.
	john := user{"John", "john@email.com"}
	john.notify()
	john.changeEmail("johnV2@gmail.com")
	john.notify()

	// Les pointeurs de type user peuvent également être utilisés pour les méthodes
	// déclarées avec un récepteur de valeur.
	emma := &user{"Emma", "emma@email.com"}
	emma.notify()
	emma.changeEmail("emmaV2@gmail.com")
	emma.notify()
}
