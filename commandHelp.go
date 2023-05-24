package main

import (
	"fmt"
)

func commandHelp(p player) error {
	commands := getCommands()
	for key, value := range commands {
		fmt.Print(key, ": ", value.name, "   ", value.description, "\n")
	}
	return nil
}
