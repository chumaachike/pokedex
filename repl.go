package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	cfg := config{}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		inputSlice := cleanInput(input)
		if len(inputSlice) == 0 {
			continue
		}
		cliCommand, ok := registerMap[inputSlice[0]]
		if !ok {
			fmt.Println("Command not found")
			continue
		}
		args := inputSlice[1:]
		if err := cliCommand.callback(&cfg, args); err != nil {
			fmt.Println("Error", args)
		}

	}
}

func cleanInput(text string) []string {
	splitString := strings.Split(text, " ")
	var stringSlice []string
	for _, word := range splitString {
		word = strings.TrimSpace(word)
		if word != "" {
			word = strings.ToLower(word)
			stringSlice = append(stringSlice, word)
		} else {
			continue
		}

	}
	return stringSlice
}
