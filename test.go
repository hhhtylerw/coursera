package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{}

type Bird struct{}

type Snake struct{}

func (a Cow) Eat() string {
	return "grass"
}

func (a Bird) Eat() string {
	return "worms"
}

func (a Snake) Eat() string {
	return "mice"
}

func (a Cow) Move() string {
	return "walk"
}

func (a Bird) Move() string {
	return "fly"
}

func (a Snake) Move() string {
	return "slither"
}

func (a Cow) Speak() string {
	return "moo"
}

func (a Bird) Speak() string {
	return "peep"
}

func (a Snake) Speak() string {
	return "hsss"
}

var m = map[string]Animal{}

func addAnimal(name string, animalType string) {
	animals := map[string]Animal{
		"cow":   Cow{},
		"bird":  Bird{},
		"snake": Snake{},
	}
	if _, ok := animals[animalType]; ok {
		m[name] = animals[animalType]
		fmt.Println("Created it!")
	} else {
		fmt.Println("Wrong type!")
	}
}

func queryByName(name string, action string) {
	funcs := map[string]func(Animal) string{
		"eat":   (Animal).Eat,
		"move":  (Animal).Move,
		"speak": (Animal).Speak,
	}
	if _, ok := funcs[action]; ok {
		fmt.Println(funcs[action](m[name]))
	} else {
		fmt.Println("Wrong action!")
	}
}

func trimNewLine(input string) string {
	switch os := runtime.GOOS; os {
	case "windows":
		// CRLF
		input = strings.Replace(input, "\r\n", "", -1)
	case "darwin":
		// LF
		input = strings.Replace(input, "\n", "", -1)
	case "linux":
		// LF
		input = strings.Replace(input, "\n", "", -1)
	default:
		// LF
		// freebsd, openbsd, plan9, ...
		input = strings.Replace(input, "\n", "", -1)
	}
	return input
}

func main() {
	funcs := map[string]func(string, string){
		"newanimal": addAnimal,
		"query":     queryByName,
	}

	for {
		consoleReader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		input, _ := consoleReader.ReadString('\n')

		input = trimNewLine(input)

		params := strings.Split(input, " ")

		if _, ok := funcs[params[0]]; ok {
			funcs[params[0]](params[1], params[2])
		} else {
			fmt.Println("Wrong command!")
		}
	}
}
