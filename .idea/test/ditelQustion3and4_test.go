package main

import (
	"strings"
	"testing"
)

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
func TestISPalindrum(t *testing.T) {
	expected := true
	actual := IsPalindrum("madam")
	if actual != expected {
		t.Errorf("Expected %t, got %t", expected, actual)
	}
}

func IsPalindrum(s string) bool {
	for count, j := 0, len(s)-1; count < j; count, j = count+1, j-1 {
		if s[count] != s[j] {
			return false
		}
	}
	return true

}

func TestToFindMinMaxAndSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 10, 55}
	min, max, sum := FindMinMaxAndSum(numbers)
	if min != expected[0] || max != expected[1] || sum != expected[2] {
		t.Errorf("Expected %d, %d, %d, got %d, %d, %d", expected[0], expected[1], expected[2], min, max, sum)
	}

}

func FindMinMaxAndSum(numbers []int) (int, int, int) {
	min := numbers[0]
	max := numbers[0]
	sum := 0
	for _, v := range numbers {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
		sum += v
	}
	return min, max, sum

}

func TestToFindAndSumNumbersFrom1To30ThatAreDivisibleBy3(t *testing.T) {
	expected := 165
	actual := FindAndSumNumbersFrom1To30ThatAreDivisibleBy3()
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func FindAndSumNumbersFrom1To30ThatAreDivisibleBy3() int {
	sum := 0
	for count := 1; count <= 30; count++ {
		if count%3 == 0 {
			sum += count
		}
	}
	return sum

}

func TestTrinagePrinting(t *testing.T) {
	expected := "   *\n  ***\n *****\n*******\n"
	actual := TrianglePrinting()
	if actual != expected {
		t.Errorf("Expected:\n%sGot:\n%s", expected, actual)
	}
}

func TrianglePrinting() string {
	var result strings.Builder

	for count := 1; count <= 4; count++ {

		for j := 1; j <= 4-count; j++ {
			result.WriteString(" ")
		}

		for j := 1; j <= 2*count-1; j++ {
			result.WriteString("*")
		}

		result.WriteString("\n")
	}

	return result.String()
}

func TestToPrintADiamoundShape(t *testing.T) {
	expected := "   *\n  ***\n *****\n*******\n *****\n  ***\n   *\n"
	actual := DiamondPrinting()
	if actual != expected {
		t.Errorf("Expected:\n%sGot:\n%s", expected, actual)
	}
}

func TestToAddTwoNumbersWIthoutPLus(t *testing.T) {
	expected := 6
	actual := SumNumber()
	if actual != expected {
		t.Errorf("Expected:\n%dGot:\n%d", expected, actual)
	}

}

func SumNumber() int {
	numbers := 5
	numbers2 := 1
	return numbers - (-numbers2)

}
func DiamondPrinting() string {
	var result strings.Builder

	for count := 1; count <= 4; count++ {

		for j := 1; j <= 4-count; j++ {
			result.WriteString(" ")
		}

		for j := 1; j <= 2*count-1; j++ {
			result.WriteString("*")
		}

		result.WriteString("\n")
	}

	for count := 3; count >= 1; count-- {

		for j := 1; j <= 4-count; j++ {
			result.WriteString(" ")
		}

		for j := 1; j <= 2*count-1; j++ {
			result.WriteString("*")
		}

		result.WriteString("\n")
	}

	return result.String()

}
