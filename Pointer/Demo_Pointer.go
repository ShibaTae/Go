package main

import "fmt"

func main() {
	msg := "Message"
	var msgPointer = &msg
	fmt.Println(*msgPointer)
	fmt.Println(&msg)

	changeMessage(&msg, "NewMassage 1")
	fmt.Println(msg)
	fmt.Println(&msg)

	changeMessage(msgPointer, "NewMassage 2")
	fmt.Println(msg)
	fmt.Println(&msg)

	changeMessage(msgPointer, "NewMassage 3")
	fmt.Println(*msgPointer)
	fmt.Println(&msg)
}

func changeMessage(aPointer *string, newMessage string) {
	*aPointer = newMessage
}
