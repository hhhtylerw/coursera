/* The goal of this activity is to explore the use of threads by creating a program for sorting integers that uses four
goroutines to create four sub-arrays and then merge the arrays into a single array.

Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted
by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge
the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should
print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {

	// Prompt for integers
	fmt.Println("")
	fmt.Println("Enter integers on a single row: ")
	var strinput string

	// Read lines to create inital array of integers
	reader := bufio.NewReader(os.Stdin)
	strinput, _ = reader.ReadString('\n')

	// Split line by spaces to create array of strings
	aryinput := strings.Split(strinput, " ")

	// Filter array to add only integers to slice
	count := 0
	var sequence []int
	for _, v := range aryinput {
		// Remove any spaces that might surround the integer
		v = strings.TrimSpace(v)
		// See if 0 entered and add it as integer to the slice
		if v == "0" {
			sequence = append(sequence, 0)
			//equence[count] = 0
			count++
		} else {
			// Convert string to integer
			valueint, _ := strconv.Atoi(v)
			// If integer is not 0, than the value of the string was an actual integer
			// Add the intger to the slice
			if valueint != 0 {
				sequence = append(sequence, valueint)
				//sequence[count] = valueint
				count++
			}
		}
	}

	// Get lengths for arrays
	totalints := len(sequence)
	avglength := totalints / 4
	fourthlength := totalints - (avglength * 3)

	// Create four subarrays of approximately equal size
	end := 0
	if fourthlength > avglength {
		end = avglength + 1
		fourthlength = fourthlength - 1
	} else {
		end = avglength
	}
	var oneArray []int = sequence[0:end]

	start := end
	if fourthlength > avglength {
		end = start + avglength + 1
		fourthlength = fourthlength - 1
	} else {
		end = start + avglength
	}
	var twoArray []int = sequence[start:end]

	start = end
	if fourthlength > avglength {
		end = start + avglength + 1
		fourthlength = fourthlength - 1
	} else {
		end = start + avglength
	}
	var threeArray []int = sequence[start:end]

	start = end
	var fourArray []int = sequence[start:]

	// Send four arrarys to four routines for sorting
	var wg sync.WaitGroup
	wg.Add(4)
	fmt.Println("")
	go subarray(oneArray, "One", &wg)
	go subarray(twoArray, "Two", &wg)
	go subarray(threeArray, "Three", &wg)
	go subarray(fourArray, "Four", &wg)
	wg.Wait()

	// Merge the 4 sorted subarrays into one large sorted array
	finalArray := append(oneArray, twoArray...)
	finalArray = append(finalArray, threeArray...)
	finalArray = append(finalArray, fourArray...)
	// Final sort
	sort.Ints(finalArray)

	fmt.Println("")

	// Print the entire sorted list
	fmt.Println("Final Sorted Array: ", finalArray)
	fmt.Println("")
}

func subarray(a []int, s string, wg *sync.WaitGroup) {
	// Print the subarray that it will sort
	fmt.Println("Subarray", s, ": ", a)
	// Sort subarray
	sort.Ints(a)
	wg.Done()
}
