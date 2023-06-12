package main

import "Golang_demo/Struct_with_constructor/employee"

func main() {
	e := employee.New(
		"Sam",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()
}
