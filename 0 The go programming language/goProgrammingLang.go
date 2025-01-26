package main

import (
	"bufio"
	"fmt"
	"os"
)

//	func main() {
//		s, sep := "", ""
//		for _, arg := range os.Args[1:] {
//			s += sep + arg
//			sep = " "
//		}
//		fmt.Println(os.Args[0])
//		for i := 1; i < len(os.Args); i++ {
//			s += sep + os.Args[i]
//			sep = " "
//		}
//		fmt.Println(s)
//	}
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
