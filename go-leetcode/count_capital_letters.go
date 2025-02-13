package main

func myCountOfCapitalLetters(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0

	for _, char_count := range s {
		if char_count >= 'A' && char_count <= 'Z' {
			count++
		}
	}
	return count
}

// func main() {
// 	inputs := []string{"GeOffrEyDuncAn", "", "12345", "OnlyUppercase"}
// 	for _, input := range inputs {
// 		count := myCountOfCapitalLetters(input)
// 		fmt.Printf("Input: %q, Number of capital letters: %d\n ", input, count)
// 	}

// }
