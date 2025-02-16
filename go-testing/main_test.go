package main

import "testing"

func TestAdd(t *testing.T) {
	// Arrange
	l, r := 2, 4
	expect := 6

	// Act
	result := Add(l, r)

	// Assert
	if expect != result {
		t.Errorf("Failed to add %v and %v. Got %v, expceted %v\n",
			l, r, result, expect)
	}
}
