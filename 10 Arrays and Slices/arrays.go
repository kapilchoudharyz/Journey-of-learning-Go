package main

import "fmt"

func main() {
	fmt.Println("This is the main function.")
	// Array.
	// Declaring an array.
	//numbers := [3]int{12, 13, 14} // first we define the length of array [3] and tell go what kind of data we want to hold inside array which is int in this case, and finally we defined the values which we want to pass in the array {12, 13, 14}
	//if we don't want to define the size of array we can do something like this -
	//numbers := [...]int{12, 13, 14}
	//fmt.Printf("Grades: %v", numbers)
	//var people [3]string
	//people[0] = "Dhan"
	//people[1] = "Yog"
	//people[2] = "Anur"
	//fmt.Printf("people : %v", people)
	////Size of the array
	//fmt.Printf("Length : %v", len(people))

	//a := [...]int{1, 2, 3}
	//b := a //When we do this we are basically creating a copy of "a".
	//We can stop this copying of data if we do this -
	//b := &a // here we are pointing to memory address of a
	//b[1] = 78
	//fmt.Println(a)
	//fmt.Println(*b) //here we are accessing the value from that memory address.

	// Slices.

	a := []int{1, 2, 3}
	b := a //In slices b will point to a without using pointer.
	b[1] = 23
	fmt.Printf("%v", a)
}
