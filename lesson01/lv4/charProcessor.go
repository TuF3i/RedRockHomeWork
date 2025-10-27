package main

import "unicode"

func safetyCheck(input string) bool {
	for i := 0; i < len(input)-1; i++ {
		item := string(input[i])
		itemPost := string(input[i+1])

		if item == "/" && itemPost == "0" {
			return false
		}
	}
	return true
}

func charProcesser(input string) []string {
	var res []string
	var current string

	for i := 0; i < len(input); i++ {
		item := input[i]
		if unicode.IsDigit(rune(item)) || rune(item) == '.' {
			current += string(item)
		} else {
			if current != "" {
				res = append(res, current)
				current = ""
			}

			if !unicode.IsSpace(rune(item)) {
				res = append(res, string(item))
			}
		}

	}

	if current != "" {
		res = append(res, current)
		current = ""
	}

	return res
}
