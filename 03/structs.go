package main

import "fmt"

//type shape interface {
//	area() int
//	perimeter() int
//}

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Printf("details %+v\n", r)
	fmt.Print(r)
	println("Area: ", r.area())
}
