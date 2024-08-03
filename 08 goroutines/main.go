package main

import "fmt"

func logger(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}
func main() {
	go logger(0)
	
}
