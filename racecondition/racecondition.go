package main

import (
	"fmt"
	"time"
)

func main() {
	favorites := make(map[string]string)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Favorite food:", favorites["food"])
	}()

	time.Sleep(2 * time.Second)
	favorites["food"] = "pizza"
	fmt.Println("Favorite food:", favorites["food"])
}

/*
	A race condition is what happens when code trys to use something that hasn't yet
	been run/solved/initialized, but is expected to be. This can happen if one portion
	of a program takes longer than expected to run and another part of the program
	tries to use the aforementioned data created by the former.
*/
