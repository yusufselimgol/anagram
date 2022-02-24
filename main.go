package main

import (
	"fmt"
)

func main() {
	var string1 string
	var string2 string
	fmt.Println("Enter First String")
	fmt.Scanf("%s", &string1)
	fmt.Println("Enter Second String")
	fmt.Scanf("%s", &string2)

	if len(string1) != len(string2) {
		fmt.Println("These strings are not anagram")
	}

	hash := make(map[string]int)

	for _, r := range string1 {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}

	for _, r := range string2 {
		j := hash[string(r)]

		if j == 0 {
			hash[string(r)] = 1
		} else {
			hash[string(r)] = j + 1
		}
	}

	var isAnagram bool = true
	for _, value := range hash {
		if value%2 != 0 {
			isAnagram = false
			break
		}

	}

	if isAnagram {
		fmt.Println("These strings are anagrams")
	} else {
		fmt.Println("These strings not are anagrams")
	}
}
