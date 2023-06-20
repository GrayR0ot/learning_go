package main

import (
	"estiam/dictionary"
	"fmt"
)

func main() {

	const nameKey string = "name"

	dict := dictionary.New()
	fmt.Printf("dict > len: %v\n", dict.Size())

	dict.Add(nameKey, "Léo")
	fmt.Printf("dict > newLen: %v\n", dict.Size())

	name, err := dict.Get(nameKey)
	if exist := err != nil; exist {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dict > name: %v\n", name)
	}

	dict.Remove(nameKey)
	name, err = dict.Get(nameKey)
	if exist := err != nil; exist {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dict > name: %v\n", name)
	}
}
