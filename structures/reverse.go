package structures

import (
	"regexp"
	"strings"
)

	var re = regexp.MustCompile(`(\s+)`)
/*
Given a string s, reverse the order of characters in each word within a sentence
while still preserving whitespace and initial word order.

Make 2 solutions here. 1 without using any go standard library (including string),
2 using any standard package and functions you want.

Compare performance and memory allocations for two solutions here.
*/

var space_chars = [3]rune {
	'\n', '	', ' ', // enter tab and space
}
func isSpace(r rune) bool {
	for _, s := range space_chars {
		if r == s {
			return true
		}
	} 
	return false
}

func reverseWord(str string) string {
	if len(str) == 0 {
		return ""
	}
	reversed := []rune(str)
	for i, j := 0, len(reversed) - 1; j > i; i, j = i+1, j-1 {
		reversed[i], reversed [j] = reversed[j], reversed[i]
	}
	return string(reversed)
}

func reverseFromTo(sl []rune, i, j int) {
	for k, l := i, j; k < l; k, l = k+1, l-1 {
		sl[k], sl[l] = sl[l], sl[k]
	}
}

func splitWithDelimiter(text string, re *regexp.Regexp) ([]string) {
	
	matches := re.FindAllStringIndex(text, -1)
	result := []string{}
	lastIndex := 0
	
	for _, match := range matches {
		start, end := match[0], match[1]
		
		if start > lastIndex {
			result = append(result, text[lastIndex:start])
		}
		
		result = append(result, text[start:end])
		lastIndex = end
	}
	
	if lastIndex < len(text) {
		result = append(result, text[lastIndex:])
	}
	
	return result
}

func Reverse1(str string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	reversed := []rune(str)
	for i, j := 0, 0; i < len(str) - 1 && j < len(str) - 1; {
		for isSpace(reversed[i]) && i < len(str) - 1 { // finding the first letter of the next word
			i++
		}
		j = i
		for !isSpace(reversed[j]) && j < len(str) - 1 { // finding the last letter of the last word
			j++
		}
		if j < len(str) - 1 { // we'll end up a char after the last word if it doesn't end the sentence
			j--
		}
		reverseFromTo(reversed, i, j)
		if j < len(str) - 1 {
			j++
			i = j
		}
	}
	return string(reversed)
}

func Reverse2(str string) string {
	parts := splitWithDelimiter(str, re)

	var result strings.Builder
	for _, part := range parts {
		if re.MatchString(part) {
			result.WriteString(part)
		} else {
			result.WriteString(reverseWord(part))
		}
	}

	return result.String()
}

