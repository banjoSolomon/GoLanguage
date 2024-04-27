package main

import "fmt"

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	fmt.Println(add(1, 2, 3))

	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")

	x := 5
	zero(&x)
	fmt.Println(x)

}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}

}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func fibonacci(x uint) uint {
	if x == 0 || x == 1 {
		return 1
	}
	return fibonacci(x-1) + fibonacci(x-2)
}

func zero(xPtr *int) {
	*xPtr = 0
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}
