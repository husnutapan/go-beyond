package main

func isUniqueCharacter(input string) bool {
	untype := make(map[rune]struct{})
	for _, r := range input {
		if _, ok := untype[r]; ok {
			return false
		} else {
			untype[r] = struct{}{}
		}
	}
	return true
}
