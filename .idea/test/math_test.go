package test

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got", v)
	}
}

func Average(float64s []float64) float64 {
	var sum float64
	for _, v := range float64s {
		sum += v
	}
	return sum / float64(len(float64s))

}
