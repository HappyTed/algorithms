// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"regexp"
	"strings"
)			

func removeNonAlphanumericChars(s string) string {
    reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	processedString := reg.ReplaceAllString(s, "")
    return processedString
}

func isPalindrome(s string) bool {
    lower := strings.ToLower(s)
	trimmed := strings.TrimSpace(lower)
	nonAlphCh := removeNonAlphanumericChars(trimmed)

	left := 0
	right := len(nonAlphCh) - 1

	for left < len(nonAlphCh) || right > 0 {
		if nonAlphCh[left] != nonAlphCh[right] {
			return  false
		}
		left++; right--
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
}