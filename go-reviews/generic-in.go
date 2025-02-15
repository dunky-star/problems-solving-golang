package main

import "fmt"

func main() {
	a1 := []int{10, 20, 30}
	// a2 := []float64{3.14, 11.7, 4.4}
	// a3 := []string{"Geo", "Dun", "Karl"}

	s1 := addVal(a1)

	fmt.Printf("Sum of %v: %v\n", a1, s1)

	testScores := []float32{
		97.7,
		100.0,
		65.5,
		29,
		30,
	}

	testScoresMap := map[string]float64{
		"Geoffrey": 97.7,
		"Amit":     98,
		"Macs":     100,
		"Ronald":   97,
		"Neville":  99,
	}

	cScores := cloneSlice(testScores)
	cMap := cloneMap(testScoresMap)

	fmt.Println(&testScores[0], &cScores[0], cScores)
	fmt.Println(cMap)
}

func cloneSlice[V any](s []V) []V {
	result := make([]V, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}

func cloneMap[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func addVal(s []int) int {
	var result int
	for _, v := range s {
		result += v
	}
	return result
}
