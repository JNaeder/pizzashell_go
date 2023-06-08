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
		fmt.Print("(> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		words := strings.Split(input, " ")

		switch words[0] {
		case "help":
			helpMessage()
		case "exit":
			fmt.Println("Goodbye!")
			os.Exit(0)
		}
	}

}

func helpMessage() {
	fmt.Print("Pizza Help\n\n")
	fmt.Print("Commands:\n")
	fmt.Print("\thelp\tshow this message\n")
	fmt.Print("\texit\texit this program\n")
}
