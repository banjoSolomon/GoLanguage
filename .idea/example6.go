package main

import "fmt"

func main() {
	element := make(map[string]string)
	element["H"] = "Hydrogen"
	element["He"] = "Helium"
	element["Li"] = "Lithium"
	element["Be"] = "Beryllium"
	element["B"] = "Boron"
	element["C"] = "Carbon"
	element["N"] = "Nitrogen"
	element["O"] = "Oxygen"
	element["F"] = "Fluorine"
	element["Ne"] = "Neon"

	fmt.Println(element["Li"])

}
