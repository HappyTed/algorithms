package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func removeAllNonAlphanumeric(s string) string {
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	clean := reg.ReplaceAllString(s, "")
	return clean
}

func isPalindrome(s string) bool {
	lower_string := strings.ToLower(s)
	prepare_string := removeAllNonAlphanumeric(lower_string)

	forward_slice := strings.Split(prepare_string, " ")

	backward_slice := strings.Split(prepare_string, " ")
	slices.Reverse(backward_slice)

	fmt.Println(forward_slice, backward_slice)

	forward := strings.Join(forward_slice, "")
	backward := strings.Join(backward_slice, "")

	fmt.Println(forward, backward)

	return forward == backward
}

var ()

func main() {
	fmt.Println(isPalindrome("race a car"))
}
