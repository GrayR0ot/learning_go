package dictionary

import (
	"errors"
	"fmt"
)

type Entry struct {
	value string
}

func (e Entry) String() string {

	return e.value
}

type Dictionary struct {
	entries map[string]Entry
}

func New() *Dictionary {

	dict := &Dictionary{
		entries: make(map[string]Entry),
	}
	return dict
}

func (d *Dictionary) Add(word string, definition string) {
	d.entries[word] = Entry{value: definition}
}

func (d *Dictionary) Get(word string) (Entry, error) {
	value, contains := d.entries[word]
	if contains {
		return value, nil
	} else {
		return Entry{}, errors.New("The key " + word + " is not in our dictionary")
	}
}

func (d *Dictionary) Remove(word string) {
	delete(d.entries, word)
	fmt.Println("Removed value for key", word, "successfully!")
}

func (d *Dictionary) List() {
	for key, element := range d.entries {
		fmt.Println(key, ":", element)
	}
}

func (d *Dictionary) Size() int {
	return len(d.entries)
}
