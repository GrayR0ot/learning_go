package main

import (
	"bufio"
	"estiam/dictionary"
	"fmt"
	"os"
	"strings"
)

func main() {

	const helpLine = "Give me a command:"
	const worldRequest = "Give me a world:"
	scanner := bufio.NewReader(os.Stdin)

	dict := dictionary.New()

	for true { // INFINITE LOOP
		fmt.Println("Enter command (ADD, DEF, REMOVE, LIST, EXIT): ")
		cmd, _ := scanner.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		cmd = strings.ToUpper(cmd)

		switch cmd {
		case "ADD":
			fmt.Println(worldRequest)
			key, _ := scanner.ReadString('\n')

			fmt.Println("Give me a definition:")
			value, _ := scanner.ReadString('\n')
			dict.Add(strings.TrimSpace(key), strings.TrimSpace(value))

		case "DEF":
			fmt.Println(worldRequest)
			key, _ := scanner.ReadString('\n')

			value, _ := dict.Get(strings.TrimSpace(key))
			fmt.Println(value)

		case "REMOVE":
			fmt.Println(worldRequest)
			key, _ := scanner.ReadString('\n')
			dict.Remove(strings.TrimSpace(key))
			fmt.Println("Removed definition for world", key)

		case "LIST":
			dict.List()

		case "EXIT":
			os.Exit(-1)

		default:
			fmt.Println("I don't understand the command:", cmd)
		}
	}
}
