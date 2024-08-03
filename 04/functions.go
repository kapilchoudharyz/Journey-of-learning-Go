package main

import (
	"errors"
	"fmt"
)

func concat(s1, s2 string) string {
	return s1 + s2
}

func main() {
	//var subtracted int = sub(1, 2)
	//fmt.Println(subtracted)
	//var added int = add(1, 2)
	//fmt.Println(added)
	//fmt.Println(concat("Hello", " world!"))
	//firstName, _ := getNames() // Basically we are ignoring the second return value
	//fmt.Println(firstName)

	//x, y := getCords()
	//fmt.Println(x, y)
	defer fmt.Println("world")
	fmt.Println("Hello")

}
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Cannot divide by zero.")
	}
	return dividend / divisor, nil
}
func getCords() (x, y int) {
	return // naked return
}

func getNames() (string, string) {
	return "kapil", "choudhary"
}
func sub(x, y int) int {
	return x - y
}
func add(x, y int) int {
	return x + y
}
