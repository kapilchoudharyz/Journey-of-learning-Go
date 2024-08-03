package main

import "fmt"

func main() {
	f(12)
	fmt.Println("Returned normally from f.")
}

func f(x int) {
	fmt.Println(x)
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	h(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Help me!")
		return
		//panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
func h(i int) {
	if i > 3 {
		fmt.Println("panicking in h")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	h(i + 1)
}
