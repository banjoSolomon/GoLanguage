package main

import "testing"

func TestFindLargestNumber1To10(t *testing.T) {
	expected := 10
	actual := FindLargestNumber1To10()
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func FindLargestNumber1To10() int {
	largest := 1
	for count := 1; count <= 10; count++ {
		if count > largest {
			largest = count
		}
	}
	return largest
}
