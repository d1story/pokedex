package main

import (
	"os"
)

func commandExit(p player) error {
	os.Exit(1)
	return nil
}
