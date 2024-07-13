package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to the Pockedex help menu!")
	fmt.Println("Here are your available commands:")

	avaliableCommands := getCommands()
	for _, cmd := range avaliableCommands {
		fmt.Printf("- %s: %s \n", cmd.name, cmd.description)
	}
	fmt.Println("")
	return nil
}
