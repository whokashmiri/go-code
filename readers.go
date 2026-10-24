package main

import (
"bufio"
 "fmt"
  "os"
  "strconv"
  "strings")

func main(){
//
// reader :=  bufio.NewReader(os.Stdin)
//
// fmt.Println("Enter name :")
//
// name , err := reader.ReadString('\n')
//
// if err != nil {
// fmt.Println("Error" , err)
// return
// }
//
// name = strings.TrimSpace(name)
// fmt.Println("Your name : " , name)



reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter asset name: ")
    nameInput, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Input error:", err)
        return
    }

    name := strings.TrimSpace(nameInput)

    fmt.Print("Enter manufacturing year: ")
    yearInput, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Input error:", err)
        return
    }

    year, err := strconv.Atoi(strings.TrimSpace(yearInput))
    if err != nil {
        fmt.Println("Invalid manufacturing year")
        return
    }

    fmt.Print("Enter estimated value (SAR): ")
    valueInput, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Input error:", err)
        return
    }

    value, err := strconv.ParseFloat(
        strings.TrimSpace(valueInput),
        64,
    )
    if err != nil {
        fmt.Println("Invalid estimated value")
        return
    }

    fmt.Println("\n--- Asset Summary ---")
    fmt.Printf("Asset: %s\n", name)
    fmt.Printf("Year: %d\n", year)
    fmt.Printf("Estimated Value: %.2f SAR\n", value)


}