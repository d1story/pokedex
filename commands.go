package main

type command struct {
	name, description string
	do                func() error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			do:          commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			do:          commandExit,
		},
	}
}
