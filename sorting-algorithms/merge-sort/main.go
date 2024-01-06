package main

import "fmt"

func merge(arr1, arr2 []int) []int {
	result := []int{}
	i := 0
	j := 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		}
		result = append(result, arr2[j])
		j++
	}

	for i < len(arr1) {
		result = append(result, arr1[i])
		i++
	}

	for j < len(arr2) {
		result = append(result, arr2[j])
		j++
	}

	return result
}

func main() {
	fmt.Println(merge([]int{2, 4, 6}, []int{1, 3, 5}))
}
