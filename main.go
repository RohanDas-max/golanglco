package main

import (
	"fmt"
)

func main() {

	//ARRAY
	var fruitlist [4]string

	fruitlist[0] = "apple"
	fruitlist[1] = "mango"
	fruitlist[2] = "cherry"
	fruitlist[3] = "berry"

	var veglist = [3]string{"potato", "Onion", "ginger"}

	fmt.Println(fruitlist)
	fmt.Println(veglist)
	fmt.Println(len(veglist))

	//!SLICES

}
