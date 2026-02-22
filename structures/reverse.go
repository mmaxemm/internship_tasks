package structures

import "strings"

/*
Given a string s, reverse the order of characters in each word within a sentence
while still preserving whitespace and initial word order.

Make 2 solutions here. 1 without using any go standard library (including string),
2 using any standard package and functions you want.

Compare performance and memory allocations for two solutions here.
*/

func Reverse1(str string) string {
	if len(str) == 0 {
		return ""
	}
	reversed := []rune(str)
	for i, j := 0, len(reversed) - 1; j > i; i, j = i+1, j-1 {
		reversed[i], reversed [j] = reversed[j], reversed[i]
	}
	return string(reversed)
}

func Reverse2(str string) string {
	runes := []rune(str)
    var sb strings.Builder
    for i := len(runes) - 1; i >= 0; i-- {
        sb.WriteRune(runes[i])
    }
    return sb.String()
}
