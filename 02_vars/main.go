package main

import "fmt"

func main() {
	//var name = "Brad"
	var age = 37
	const isCool = true
	var size float32 = 2.3

	//shorthand only inside functions
	// name := "Brad"
	// email := "brad@gmail.com"

	name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
