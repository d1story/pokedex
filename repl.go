package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl(p player) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		reader.Scan()
		text := cleanText(reader.Text())
		err := getCommands()[text[0]].do(p)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func cleanText(a string) []string {
	a = strings.ToLower(a)
	b := strings.Fields(a)
	return b
}
