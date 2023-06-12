package main

import "fmt"

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("", numbers)
	fmt.Println("", numbers["three"])

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#0ff"
	fmt.Println("", colors)
	fmt.Println("", colors["green"])
	fmt.Println("", colors["blue"])

	var courses = make(map[string]map[string]int)
	courses["Android"] = make(map[string]int)
	courses["Android"]["price"] = 200
	courses["Android"]["code"] = 1234

	courses["iOS"] = make(map[string]int)
	courses["iOS"]["price"] = 500
	courses["iOS"]["code"] = 4321
	fmt.Println(courses)
}
