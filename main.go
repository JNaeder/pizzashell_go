package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Welcome Message
	fmt.Println("Welcome to Pizza Shell!")

	//Main Loop
	for {
		words := getUserInput("(> ")

		switch words[0] {
		case "help":
			helpMessage()
		case "exit":
			fmt.Println("Goodbye!")
			os.Exit(0)
		case "get":
			//----- GET COMMANDS -----
			if len(words) < 2 {
				getHelp()
				break
			}
			switch words[1] {
			case "location":
				fmt.Println("Your Location ->")
				fmt.Println(GetIp())
			case "test":
				for {
					testInput := getUserInput("&&> ")
					fmt.Println(testInput)
					if testInput[0] == "exit" {
						break
					}
				}
			}
		default:
			fmt.Printf("'%s' is not a recognized command. Type 'help' for a list of commands.\n", words[0])
		}
	}

}

func getUserInput(prompt string) []string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	words := strings.Split(input, " ")
	return words
}

func helpMessage() {
	fmt.Print("Pizza Help\n\n")
	fmt.Print("Commands:\n")
	fmt.Print("\thelp\tshow this message\n")
	fmt.Print("\texit\texit this program\n")
}

func getHelp() {
	fmt.Print("No command recognized after 'get'. Try these options:\n\n")
}
