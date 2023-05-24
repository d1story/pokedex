package main

type command struct {
	name, description string
	do                func(player) error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			do:          commandHelp,
		},
		"smf": {
			name:        "showMapF",
			description: "Displays the next 20 locations",
			do:          commandMapF,
		},
		"smb": {
			name:        "showMapB",
			description: "Displays the previous 20 locations",
			do:          commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			do:          commandExit,
		},
	}
}
