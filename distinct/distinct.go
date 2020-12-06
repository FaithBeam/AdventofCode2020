package distinct

import "strings"

func Letters(input []string) (m map[string]string) {
	m = make(map[string]string)
	for _, letter := range input {
		m[letter] = letter
	}
	return
}

func GroupLetters(input []string) (l []map[string]string) {
	for _, group := range input {
		groupSplit := strings.Split(group, "")
		l = append(l, Letters(groupSplit))
	}
	return
}
