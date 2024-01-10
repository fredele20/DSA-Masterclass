package main

import "fmt"

func findValue(nums []int, val int) int {
	for i:=0; i < len(nums); i++ {
		if nums[i] == val {
			return i
		}
	}
	return -1
}

func main() {

	arr := []int{1, 3, 5, 2, 8, 4, 9, 7, 10}
	fmt.Println(findValue(arr, 4))
	
}
