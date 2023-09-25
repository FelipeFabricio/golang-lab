package main

import (
	"os"

	"github.com/felipefabricio/golang-lab/fileHandler"
)

func main() {
	filename := os.Args[1]
	if filename == "" {
		panic("Filename is required!")
	}
	fileHandler.CreateFile(filename)

	args := os.Args[2:]
	if len(args) > 0 {
		fileHandler.WriteFile(filename, args)
	}

	fileHandler.ReadFile(filename)
	// fileHandler.DeleteAllLines(filename)
	// fileHandler.VerifyIfFileExists(filename)
	// fileHandler.ReadFile(filename)
}
