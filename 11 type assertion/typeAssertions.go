package main

import "fmt"

func main() {
	name := interface{}("Hello")
	typeOfName, ok := name.(int)
	fmt.Println(typeOfName, ok)
}
