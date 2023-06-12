package main

import "Golang_demo/Struct_with_constructor_and_pointer/employee"

func main() {
	e := employee.Init(
		"Sam",
		"Adolf",
		30,
		20)

	e.LeavesRemaining()

	e = employee.Init(
		"Tae",
		"mj",
		30,
		20)

	e.LeavesRemaining()
}
