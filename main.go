package main

import "fmt"

func main() {
	fmt.Println("String: hello", IsStringCharsUnique("hello"))
	fmt.Println("String: dog", IsStringCharsUnique("dog"))
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
