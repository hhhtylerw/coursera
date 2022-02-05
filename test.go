package main

import (
	"fmt"
	"time"
)

/*
	incrementX goroutine tries to increment X by 1
*/
func incrementX(x *int) {
	// simulating a delay (can be any computing work, or context switch)
	time.Sleep(10 * time.Millisecond)
	*x = *x + 1
}

/*
	printX goroutine tries to print X
*/
func printX(x *int) {

	fmt.Printf("After incrementing : %d", *x)
}

func main() {
	var x int = 0
	/*
		A goroutine is started on increment X. As a programmer
		my ideal output should be 1, as that is the intention of my code.
		But since both incrementX and printX goroutines are running concurrently
		The output can be 0 or 1.
		Why? RACE CONDITION! Both printX and incrementX are racing for Variable X
		There is communication between them.
			1. incrementX tries to write to X
			2. printX tries to read X
		The order of these instructions cannot be predicted and hence the unpredictable output

	*/
	go incrementX(&x)
	go printX(&x)
	time.Sleep(100 * time.Millisecond)
}
