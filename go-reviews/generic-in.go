package main

import "fmt"

func main() {
	testScores := []float32{
		97.7,
		100.0,
		65.5,
		29,
		30,
	}

	c := clone(testScores)

	fmt.Println(&testScores[0], &c[0], c)
}

func clone[V any](s []V) []V {
	result := make([]V, len(s))
	for i, v := range s {
		result[i] = v
	}
	return result
}
