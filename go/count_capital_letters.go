package main

func myCountOfCapitalLetters(s string) int {
	count := 0

	for _, char_count := range s {
		if char_count >= 'A' && char_count <= 'Z' {
			count++
		}
	}
	return count
}
