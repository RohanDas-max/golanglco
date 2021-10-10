package main

import "fmt"

func main() {
	tree := []string{"bot", "kochu"}

	for i, j := range tree {
		fmt.Println(i, j)
		fmt.Printf("%v, %v\n", i, j)
	}
}
