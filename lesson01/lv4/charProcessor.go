package main

import "unicode"

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
