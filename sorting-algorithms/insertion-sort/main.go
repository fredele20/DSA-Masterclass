package main

import "fmt"


func insertionSort(arr []int) []int {

	for i:=1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}

	return arr
}

func main() {
	fmt.Println(insertionSort([]int{2, 5, 3, 6, 9, 1, 4}))
}
