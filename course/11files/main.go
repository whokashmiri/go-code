package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files")
	content := "this needs to go in a file - learnGo.dev"

	file, err := os.Create("./content.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	defer file.Close()
	readFile("./content.txt")

}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(dataByte))
}
