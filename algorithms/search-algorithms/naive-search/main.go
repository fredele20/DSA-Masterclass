package main

import "fmt"

func main() {
	fmt.Println(naiveStringSearch("my word", "or"))
}

func naiveStringSearch(word, pattern string) int {
	count := 0

	for i := 0; i < len(word)-1; i++ {
		for j := 0; j < len(pattern); j++ {
			if word[i + j] != pattern[j] {
				break
			}
			if j == (len(pattern) - 1) {
				count++
			}
		}
	}

	return count
}
