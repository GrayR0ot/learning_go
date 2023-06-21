package repository

import (
	"dictionary/model"
	"errors"
	"fmt"
	"time"
)

type Repository struct {
	Database map[string]model.Entry
}

func (d *Repository) Add(word string, definition string) {
	d.Database[word] = model.Entry{Value: definition, Date: time.Now()}
}

func (d *Repository) Get(word string) (model.Entry, error) {
	value, contains := d.Database[word]
	if contains {
		return value, nil
	} else {
		return model.Entry{}, errors.New("The key " + word + " is not in our dictionary")
	}
}

func (d *Repository) Remove(word string) {
	delete(d.Database, word)
	fmt.Println("Removed value for key", word, "successfully!")
}

func (d *Repository) List() {
	fmt.Println("Word	| Definition			| Date")
	for key, element := range d.Database {
		fmt.Println(key, "	| ", element.Value, "		| "+element.Date.Format(time.DateTime))
	}
}

func (d *Repository) Size() int {
	return len(d.Database)
}
