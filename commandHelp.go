package main

import (
	"fmt"
)

func commandHelp() error {
	commands := getCommands()
	for key, value := range commands {
		fmt.Print(key, ": ", value.name, "   ", value.description, "\n")
	}
	return nil
}
