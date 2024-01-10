package main

import "fmt"

// Binary Search algorithm
// Binary search algorithm uses the dived and conquer algorithm pattern
// 1. Create a function, the function should accept a sorted array or slice and a value
// 2. Create a left pointer at the start of the slice, and a right pointer at the end of the array
// 3. Loop while the left pointer is lesser than the right
// 		a. Create a pointer in the middle
// 		b. Compare the middle value
//    c. If you find the value you want, return the index
//    d. If the value is too small, move the left pointer up
//    e. If the value is too big, move the right pointer down
// 4. If the value is not found, return -1

func binarySearch(arr []int, val int) int {
	if len(arr) == 0 {
		return -1
	}

	start := 0;
	end := len(arr) -1
	
	for start <= end {
		middle := (start + end) / 2
		if arr[middle] == val {
			return middle
		} else if arr[middle] > val {
			end = middle-1
		} else {
			start = middle + 1
		}
	}
	return -1
}


func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9,}

	fmt.Println(binarySearch(arr, 8))
	
}
