package main

import "fmt"

func main() {
	fmt.Println("String: hello", IsStringCharsUnique("hello"))
	fmt.Println("String: dog", IsStringCharsUnique("dog"))

	fmt.Println("String: qqwweerrt", IsPermutationOfPalindrome("qqwweerrt"))
	fmt.Println("String: qqwweerrta", IsPermutationOfPalindrome("qqwweerrta"))
}

func IsStringCharsUnique(input string) bool {
	charMap := map[string]bool{}
	charMap[string(input[0])] = true

	for i := 1; i < len(input); i++ {
		if _, exists := charMap[string(input[i])]; exists {
			return false
		} else {
			charMap[string(input[i])] = true
		}
	}

	return true
}

func IsPermutationOfPalindrome(input string) bool {
	charMap := map[string]int{}

	for _, v := range input {
		if _, exists := charMap[string(v)]; exists {
			charMap[string(v)]++
		} else {
			charMap[string(v)] = 1
		}
	}

	foundOdd := false
	for _, value := range charMap {
		if value%2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}

	return true
}
