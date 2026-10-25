package main

import "fmt"
func greet(){
fmt.Println("Welcome to functions")
}

func printAsset(name string , year int){
fmt.Println("Asset " , name)
fmt.Println("age " , year)
}


func calAge(mYear int) int {
currentYear := 2026

return currentYear - mYear
}

func main(){
// greet()

// printAsset("Camry" , 2025)

age:= calAge(2020)

fmt.Println(age)

}