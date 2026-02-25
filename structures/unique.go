package structures
/* 
For a given string slice which can have multiple repeated values create set of unique values using map
and remove all repeated values from slice respecting the original order.
*/

func Unique(sl []string) []string {
	containedWords := make(map[string]struct{})
	result := make([]string, 0, len(sl))
	num := len(containedWords)
	for _, word := range sl {
		containedWords[word] = struct{}{}
		if num < len(containedWords) {
			result = append(result, word)
			num++
		}
	}
	return result
}
