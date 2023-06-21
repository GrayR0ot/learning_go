package main

import (
	"bufio"
	"dictionary/repository"
	"fmt"
	"os"
	"strings"
)

type CLI struct {
	Repository *repository.Repository
}

func (CLI *CLI) Listen() {
	const helpLine = "Give me a command:"
	const wordRequest = "Give me a word:"
	scanner := bufio.NewReader(os.Stdin)

	for true { // INFINITE LOOP
		fmt.Println("Enter command (ADD, DEF, REMOVE, LIST, EXIT): ")
		cmd, _ := scanner.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		cmd = strings.ToUpper(cmd)

		switch cmd {
		case "ADD":
			fmt.Println(wordRequest)
			key, _ := scanner.ReadString('\n')

			fmt.Println("Give me a definition:")
			value, _ := scanner.ReadString('\n')
			CLI.Repository.Add(strings.TrimSpace(key), strings.TrimSpace(value))

		case "DEF":
			fmt.Println(wordRequest)
			key, _ := scanner.ReadString('\n')

			value, _ := CLI.Repository.Get(strings.TrimSpace(key))
			fmt.Println(value)

		case "REMOVE":
			fmt.Println(wordRequest)
			key, _ := scanner.ReadString('\n')
			CLI.Repository.Remove(strings.TrimSpace(key))
			fmt.Println("Removed definition for word", key)

		case "LIST":
			CLI.Repository.List()

		case "EXIT":
			os.Exit(-1)

		default:
			fmt.Println("I don't understand the command:", cmd)
		}
	}
}
