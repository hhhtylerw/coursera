package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg sync.WaitGroup
var ch chan []int

func main() {
	ch = make(chan []int, 4)
	arr := make([]int, 0, 12)
	out := make([]int, 0, 12)

	fmt.Println("Enter nums for array: ")
	for i := 0; i < 12; i++ {
		var num int
		fmt.Print("> ")
		fmt.Scan(&num)
		arr = append(arr, num)
	}

	for i := 0; i < len(arr); i += 4 {
		wg.Add(1)
		go Sorter(arr[i : i+4])
	}

	wg.Wait()

	for i := 0; i < len(arr); i += 4 {
		out = append(out, <-ch...)
	}
	sort.Ints(out)
	fmt.Println("After sorting:", out)
}

func Sorter(arr []int) {
	fmt.Println("Before sorting:", arr)
	sort.Ints(arr)
	ch <- arr
	wg.Done()
}
