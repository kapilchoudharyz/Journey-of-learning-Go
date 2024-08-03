package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)
func  Hello(name string) (string, error){
	// fmt.Println("Hello, How are you!")
	if name == ""{
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
func Bye(name string) (string, error){
	// fmt.Println("Hello, How are you!")
	if name == ""{
		return "", errors.New("Not right arguments")
	}
	message := fmt.Sprintf("Bye, %v. Take care!", name)
	return message, nil
}

func RandomFormat() string {
	format := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return format[rand.Intn(len(format))]
}