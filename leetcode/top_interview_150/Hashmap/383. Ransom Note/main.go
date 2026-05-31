package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	magazMap := make(map[rune]int)
	for _, let := range magazine {
		magazMap[let] += 1
	}

	for _, let := range ransomNote {
		if count, ok := magazMap[let]; ok && count > 0 {
			magazMap[let] -= 1
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}
