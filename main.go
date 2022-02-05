package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	animals := map[string]*Animal{
		"cow": {
			"grass", "walk", "moo",
		},
		"bird": {
			"worms", "fly", "peep",
		},
		"snake": {
			"mice", "slither", "hsss",
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println(">")

		scanner.Scan()
		userInput := strings.Split(scanner.Text(), " ")

		if a, exist := animals[userInput[0]]; exist {
			switch strings.ToLower(userInput[1]) {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				fmt.Printf("The %s command doesn't exist. Choose one of eat, move, speak.\n", userInput[1])
			}
		} else {
			fmt.Printf("The %s doesn't exist, please try agin\n", userInput[0])
		}

	}

}
