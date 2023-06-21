//Interfaces

// Programme d'exemple montrant comment utiliser une interface en Go. Dans ce cas,
// un pointeur est utilisé pour prendre en charge l'appel de l'interface.
package main

import (
	"fmt"
)

// notifier est une interface qui définit un comportement de type de notification.
type notifier interface {
	notify()
}

// user définit un utilisateur dans le programme.
type user struct {
	name  string
	email string
}

type group struct {
	id      int
	name    string
	members []*user
}

func (g *group) AddMember(u *user) {
	g.members = append(g.members, u)
}

// notify implémente une méthode qui peut être appelée via
// une valeur de type user.
func (u user) notify() {
	fmt.Printf("utilisateur : Envoi de l'e-mail utilisateur à %s<%s>\n",
		u.name,
		u.email)
}

func (g group) notify() {
	fmt.Printf("There is currently %v members in the group!\n", len(g.members))
}

func (g *group) UpdateMemberName(index int, newName string) {
	user := &g.members[index]
	(*user).UpdateUserName(newName)
}

func (u *user) UpdateUserName(newName string) {
	u.name = newName
}

func (g *group) GetMembers() {
	for i, user := range g.members {
		fmt.Println("[", i, "] Name: ", user.name, "Email:", user.email)
	}
}

// main est le point d'entrée de l'application.
func main() {
	// Crée une valeur de type user.
	u := user{"Jill", "jill@email.com"}

	// Passe un pointeur de type user à la fonction.
	sendNotification(u)
	u.UpdateUserName("Pierre")
	sendNotification(u)

	g := group{id: 1, name: "Admins", members: make([]*user, 0)}
	g.AddMember(&u)
	sendNotification(g)

	g.UpdateMemberName(0, "Léo")
	g.GetMembers()
	sendNotification(u)

	u.UpdateUserName("Jill")
	g.GetMembers()
}

// sendNotification accepte des valeurs qui implémentent l'interface notifier
// et envoie des notifications.
func sendNotification(notify notifier) {
	notify.notify()
}
