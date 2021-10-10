package main

func main() {

	//? Pointer
	// a := 10
	// pt := &a
	// fmt.Println(pt)
	// fmt.Println(*pt)

	//?ARRAY
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

	// var langs = []string{"js", "py", "java", "go", "ts"}
	// fmt.Println(langs)
	// var index int = 2
	// langs = append(langs[:index], langs[index+1:]...)
	// fmt.Println(langs)

	//? MAPS

	// 	langs := make(map[string]string)

	// 	langs["js"] = "javascript"
	// 	langs["py"] = "python"
	// 	langs["go"] = "golang"
	// 	langs["rb"] = "ruby"

	// 	fmt.Println(langs)
	// 	fmt.Println("js short for: ", langs["js"])
	// 	delete(langs, "rb")
	// 	fmt.Println(langs)

	// 	for key, value := range langs {
	// 		fmt.Printf("for key %v, value is %v \n", key, value)
	// 	}

	//?struct

	// rohan := User{Name: "Rohan", Email: "rohan@rohan.dev", Age: 24}
	// fmt.Printf("%v age is %v and Email is %v\n", rohan.Name, rohan.Age, rohan.Email)

	//? Switch Case

	// rand.Seed(time.Now().UnixNano())
	// dicenumber := rand.Intn(6) + 1

	// switch dicenumber {
	// case 1:

	// 	fmt.Println("The number is 1 you can open")

	// case 2:

	// 	fmt.Println("you can move 2 spot")

	// case 3:

	// 	fmt.Println("you can move 3 spot")

	// case 4:

	// 	fmt.Println("you can move 4 spot")

	// case 5:

	// 	fmt.Println("you can move 5 spot")
	// 	// fallthrough
	// case 6:
	// 	fmt.Println("the number is 6 you can roll again")

	// }

	//? For Loop
	// langs := []string{"js", "go", "py", "rb", "ts"}

	// for i := range langs {
	// 	fmt.Println(langs[i])
	// }

	// for index, i := range langs {
	// 	fmt.Printf("index is %v and language is %v\n", index, i)
	// }

	// 	value := 1

	// 	for value < 10 {

	// 		if value == 5 {
	// 			// fmt.Println("value is:", value)
	// 			goto hola
	// 			// value++
	// 			// continue
	// 		}
	// 		fmt.Println("value is ", value)
	// 		value++
	// 		// goto hola

	// 	}

	// hola:
	// 	fmt.Println("hola amigos")

	//?methods

}

// type User struct {
// 	Name   string
// 	Email  string
// 	status bool
// 	Age    int8
// }
