package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type AnimalStruct struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	// Define variables
	m := make(map[string]Animal)
	var cow Animal = AnimalStruct{"grass", "walk", "moo"}
	var bird Animal = AnimalStruct{"worms", "fly", "peep"}
	var snake Animal = AnimalStruct{"mice", "slither", "hsss"}

	for {
		// Get user input
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		inputs := strings.Split(input, " ")

		// Read input and act accordingly
		switch inputs[0] {
		case "newanimal":
			name := inputs[1]
			switch inputs[2] {
			case "cow":
				m[name] = cow
			case "bird":
				m[name] = bird
			case "snake":
				m[name] = snake
			default:
				continue
			}
			fmt.Println("Created it!")
		case "query":
			name := inputs[1]
			switch inputs[2] {
			case "eat":
				m[name].Eat()
			case "move":
				m[name].Move()
			case "speak":
				m[name].Speak()
			}

		}

	}
}

func (animal AnimalStruct) Eat() {
	fmt.Println(animal.food)
}
func (animal AnimalStruct) Move() {
	fmt.Println(animal.locomotion)
}
func (animal AnimalStruct) Speak() {
	fmt.Println(animal.noise)
}
