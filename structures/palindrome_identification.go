package structures

import (
	"regexp"
	"strings"
)

/*
A palindrome is a word, number, phrase, or other sequence of symbols that reads the same backwards as forwards,
such as the words civic or radar or numbers such as 22\2\22 or sentence such as “Mr. Owl ate my metal worm”.
Implement 3 possible versions of a program to identify if string is a polindrome,
compare performance (benchmarks), memory allocated (optional) and O(n) of each solution.
*/

func usify(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`) // "not a letter"
	return strings.ToLower(re.ReplaceAllString(s, ""))
}

func IsPalindrome1(str string) bool {
	str = usify(str)
	if len(str) == 0 {
		return true
	}
	for i, j := 0, len(str) - 1; j > i; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func IsPalindrome2(str string) bool {
	str = usify(str)
	reverse := []rune(str)
    for i, j := 0, len(reverse)-1; i < j; i, j = i+1, j-1 {
        reverse[i], reverse[j] = reverse[j], reverse[i]
    }
    return string(reverse) == str
}

func IsPalindrome3(str string) bool {
	str = usify(str)
	runes := []rune(str)
    
    var recursiveCheck func(int, int) bool
    recursiveCheck = func(i, j int) bool {
        if i >= j {
            return true
        }
        if runes[i] != runes[j] {
            return false
        }
        return recursiveCheck(i+1, j-1)
    }
    
    return recursiveCheck(0, len(runes)-1)
}
