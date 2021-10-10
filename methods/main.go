package main

import "fmt"

func main() {
	user := User{"Rohan", "rohan@rohan.dev", true, 26}

	// fmt.Println(user)
	// fmt.Printf("%+v\n", user)

	user.GetStatus()
	user.NewEmail()

	fmt.Println("original email is: ", user.Email)
}

type User struct {
	Name   string
	Email  string
	status bool
	Age    int8
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.status)
}

func (u User) NewEmail() {
	u.Email = "rohang@test.dev"
	fmt.Println("new copy of email is: ", u.Email)
}
