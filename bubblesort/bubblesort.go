package main

import "fmt"

func main() {
	nums := make([]int, 10)
	fmt.Println("Enter 10 numbers:")
	for i := 0; i < 10; i++ {
		var input int
		fmt.Print("> ")
		fmt.Scan(&input)
		nums[i] = input
	}

	fmt.Println("Before sorting:", nums)
	BubbleSort(nums)
	fmt.Println("After sorting:", nums)
}

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)
			}
		}
	}
}

func Swap(nums []int, index int) {
	temp := nums[index]
	nums[index] = nums[index+1]
	nums[index+1] = temp
}
