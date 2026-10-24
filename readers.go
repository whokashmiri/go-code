package main

import (
"bufio"
 "fmt"
  "os"
  "strings")

func main(){

reader :=  bufio.NewReader(os.Stdin)

fmt.Println("Enter name :")

name , err := reader.ReadString('\n')

if err != nil {
fmt.Println("Error" , err)
return
}

name = strings.TrimSpace(name)
fmt.Println("Your name : " , name)

}