package main

import "fmt"

type product struct {
	name  string
	price int
	stock int
}

func main() {
	var p1 product
	p1.name = "Pepsi"
	p1.price = 15
	p1.stock = 100
	show(p1)
}

func show(p product) {
	fmt.Println(p)
}
