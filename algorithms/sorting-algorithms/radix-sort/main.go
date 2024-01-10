package main

import (
	"fmt"
	"math"
	"strconv"
)

func getDigit(num, place int) int {
	return (num / int(math.Pow(10, float64(place)))) % 10
}

func digitCount(num int) int {
	// return int(math.Floor(math.Log10(math.Abs(float64(num))))) + 1
	return len(strconv.Itoa(num))
}

func mostDigit(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		if digCount := digitCount(nums[i]); digCount > count {
			count = digCount
		}
	}

	return count
}

func radixSort(nums []int){
	highestDigit := mostDigit(nums)

	for i := 0; i <= highestDigit; i++ {
		arr := make([]int, 10)
		// arr1 := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			digit := getDigit(nums[j], i)
			arr[digit] = nums[j]
			fmt.Println(arr)
		}
	}
}

func main() {

	fmt.Println(getDigit(1234587, 10))
	fmt.Println(digitCount(0))

	fmt.Println(mostDigit([]int{2398343094, 1, 657, 1234, 9890, 673, 894839}))
	radixSort([]int{2398343094, 1, 657, 1234, 9890, 673, 894839})

}
