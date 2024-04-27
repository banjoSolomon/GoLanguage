package main

import "fmt"

func main() {
	for count := 1; count <= 10; count++ {
		if count%2 == 0 {
			fmt.Println(count, "even")
		} else {
			fmt.Println(count, "odd")
		}
	}

}
