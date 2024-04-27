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

func TestTOFindTwoLargestNumbers(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{10, 9}
	largest1, largest2 := FindTwoLargestNumber(numbers)
	if largest1 != expected[0] || largest2 != expected[1] {
		t.Errorf("Expected %d, %d, got %d, %d", expected[0], expected[1], largest1, largest2)
	}

}

func FindTwoLargestNumber(numbers []int) (int, int) {
	largest1 := numbers[0]
	largest2 := numbers[0]
	for _, v := range numbers {
		if v > largest1 {
			largest2 = largest1
			largest1 = v
		} else if v > largest2 {
			largest2 = v
		}
	}
	return largest1, largest2

}
