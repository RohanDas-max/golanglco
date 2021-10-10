package main

import "fmt"

func main() {

	defer fmt.Println("World")
	fmt.Println("Hello")
	randNum()
}

func randNum() {

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
