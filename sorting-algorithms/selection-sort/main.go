package main

import "fmt"

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if i != min {
			temp := arr[i]
			arr[i] = arr[min]
			arr[min] = temp
		}
	}
	return arr
}

func main() {
	fmt.Println(selectionSort([]int{1, 2, 5, 3, 1, 4, 6}))
}
