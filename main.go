package main

import (
	"fmt"
)

func main() {

	//ARRAY
	// var fruitlist [4]string

	// fruitlist[0] = "apple"
	// fruitlist[1] = "mango"
	// fruitlist[2] = "cherry"
	// fruitlist[3] = "berry"

	// var veglist = [3]string{"potato", "Onion", "ginger"}

	// fmt.Println(fruitlist)
	// fmt.Println(veglist)
	// fmt.Println(len(veglist))

	//!SLICES

	// var Vegies = []string{"potato", "onion", "garlic"}
	// fmt.Println(Vegies)
	// Vegies = append(Vegies, "ginger", "cabbage")
	// fmt.Println(Vegies)
	// Vegies = append(Vegies[1:3])
	// fmt.Println(Vegies)

	//?adding making sorting
	// highscores := make([]int, 4)
	// highscores[0] = 999
	// highscores[1] = 231
	// highscores[2] = 331
	// highscores[3] = 31

	// fmt.Println(highscores)

	// highscores = append(highscores, 432, 65)
	// fmt.Println(highscores)
	// fmt.Println(len(highscores))

	// sort.Ints(highscores)
	// fmt.Println(highscores)
	// fmt.Println(sort.IntsAreSorted(highscores))

	//?remove value from slice based on index

	var langs = []string{"js", "py", "java", "go", "ts"}
	fmt.Println(langs)
	var index int = 2
	langs = append(langs[:index], langs[index+1:]...)
	fmt.Println(langs)

}
