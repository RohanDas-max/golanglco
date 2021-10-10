package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	content := "whassssaaaaauuuuupppp"

	file, err := os.Create("./myfile.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("length is :", length)
	defer file.Close()
	readFile("/home/rohan/work/go/golanglco/fileshandling/myfile.txt")

}
func readFile(filename string) {

	data, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println(string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
