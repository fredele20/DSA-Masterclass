package main

import "fmt"

func partition(arr []int, start, end int) int {
	count := start
	pivotIndex := arr[start]

	for i := start + 1; i < len(arr); i++ {
		if pivotIndex > arr[i] {
			count++
			arr[count], arr[i] = arr[i], arr[count]
		}
	}

	arr[start], arr[count] = arr[count], arr[start]

	return count

}

func quickSort(arr []int, start, end int) []int {
	if start < end {
		fmt.Println(start, end)
		pivotIndex := partition(arr, start, end)

		quickSort(arr, start, pivotIndex-1)
		quickSort(arr, pivotIndex+1, end)
	}

	return arr
}

func main() {
	arr1 := []int{10, 100, 4, 8, 2, -1, 1, 5, 7, 6, -2, 3}
	n := len(arr1)
	fmt.Println(quickSort(arr1, 0, n-1))

}
