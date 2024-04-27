package main

import "fmt"

func main() {
	fmt.Println("Welcome to Tax Calculator Application")
	for {
		var name string
		var earnings, task float64
		fmt.Println("Enter citizen's name (or type 'quit' to exit)): ")
		fmt.Scanln(&name)
		if name == "quit" {
			break
		}
		fmt.Println("Enter earnings for the year: ")
		fmt.Scanln(&earnings)

		if earnings <= 30000 {
			task = earnings * 0.15
		} else {
			task = 30000*0.15 + (earnings-30000)*0.20
		}
		fmt.Println(name+": Total Task = $", +task)
	}

}
