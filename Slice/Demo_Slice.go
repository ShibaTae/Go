package main

import "fmt"

func main() {
	var numbers1 = make([]int, 0, 5)
	numbers1 = append(numbers1, 1)
	numbers1 = append(numbers1, 2)
	showSlice(numbers1) //len = 2 cap = 5 slice = [1 2] จองไว้ 5

	var numbers2 []int
	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	showSlice(numbers2) //len = 2 cap = 2 slice = [1 2] มีแค่ไหนจองเท่านั้น
}
func showSlice(x []int) {
	fmt.Printf("len = %d cap = %d slice = %v\n", len(x), cap(x), x)
}
