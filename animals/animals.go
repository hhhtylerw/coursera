package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func main() {
	// Declare animals
	var cow = Animal{"grass", "walk", "moo"}
	var bird = Animal{"worms", "fly", "peep"}
	var snake = Animal{"mice", "slither", "hsss"}

	for {
		// Get user input
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		inputs := strings.Split(input, " ")

		// Find animal from input
		var animal Animal
		switch inputs[0] {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			continue
		}

		// Execute command from input
		switch inputs[1] {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			continue
		}
	}
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}
