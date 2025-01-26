// package main
//
// import (
//
//	"fmt"
//	"sync"
//	"time"
//
// )
//
//	func doSomething(s string, wg *sync.WaitGroup) {
//		defer wg.Done() //way to tell wg that this go routine is executed decrement the internal counter.
//		for i := 0; i < 5; i++ {
//			time.Sleep(100 * time.Millisecond)
//			fmt.Printf("%v: %v \n", s, i)
//		}
//	}
//
//	func main() {
//		//Goroutine is a lightweight thread managed by go runtime.
//		var wg sync.WaitGroup
//		fmt.Println(wg)
//		wg.Add(1)                        //A way to tell wg that increase the count of go routines running.
//		go doSomething("Something", &wg) //We are passing the memory address so that it can work properly otherwise we will be decrementing the counter of copy in "say" function.
//		fmt.Println("Do")
//		//Doing this because function will exit immediately after printing "Hello"
//		//time.Sleep(600 * time.Millisecond)
//		wg.Wait() //Waits until the counter goes to zero ensuring that all the go routines are executed.
//	}
package main

import (
	"fmt"
	"time"
)

func doSomething(aChannel chan bool) {
	initialChannelValue := <-aChannel
	fmt.Println(initialChannelValue)
	time.Sleep(time.Second * 1)
	aChannel <- true
}

func main() {
	done := make(chan bool)
	go doSomething(done)
	done <- false
	isTrue := <-done
	fmt.Println(isTrue)
}
