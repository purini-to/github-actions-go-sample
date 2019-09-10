package main

import "fmt"

// GetMessage get sample message.
func GetMessage() string {
	return "Hello Actions!"
}

func main() {
	msg := GetMessage()
	fmt.Printf("msg: %s", msg)
}
