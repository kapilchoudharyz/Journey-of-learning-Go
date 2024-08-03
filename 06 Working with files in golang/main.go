package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Here we will be learning about handling files in golang.")
	content := "Some content to write to the file."
	file, err := os.Create("test-file.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes to the file.\n", length)
	err = file.Close()
	if err != nil {
		return
	}
	readFromFile()

}

func readFromFile() {
	file, err := os.Open("test-file.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		panic(err)
	}
	str := string(bs)
	fmt.Println(str)
}
