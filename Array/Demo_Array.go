package main

import "fmt"

func main() {
	var array1 = [3]int{1, 2, 3}

	fmt.Print("\n")
	for _, item := range array1 {
		fmt.Print(item, ",")
	}
}
