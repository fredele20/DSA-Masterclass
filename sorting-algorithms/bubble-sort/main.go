package main

import "fmt"


func bubbleSort(arr []int) []int {
	var noSwaps bool
	for i := len(arr); i > 0; i-- {
		noSwaps = true
		for j := 0; j < i - 1; j++ {
			if arr[j] > arr[j + 1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				noSwaps = false
			}
		}
		if noSwaps {
			break
		}
	}
	return arr
}

func main() {

	fmt.Println(bubbleSort([]int{12, 5, 37, 25, -3, 88, 92, 14}))
	
}
