package main

import "fmt"

func mainGen() {
	a1 := []int{10, 20, 30}
	a2 := []float64{3.14, 11.7, 4.4}
	a3 := []string{"Geo", "Dun", "Karl"}

	s1 := addVal(a1)
	s2 := addVal(a2)
	s3 := addVal(a3)

	fmt.Printf("Sum of %v: %v\n", a1, s1)
	fmt.Printf("Sum of %v: %v\n", a2, s2)
	fmt.Printf("Sum of %v: %v\n", a3, s3)

	// Slice
	testScores := []float32{
		97.7,
		100.0,
		65.5,
		29,
		30,
	}

	// Map
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

// Define our own Generic with interface
type addable interface {
	int | float64 | string
}

func addVal[V addable](s []V) V {
	var result V
	for _, v := range s {
		result += v
	}
	return result
}
