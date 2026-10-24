
package main

import (
    "fmt"
    "strconv"
    "bufio"
    "os"
    "strings"
)

func main() {
//     yearInput := "2023"
//
//    year , err  := strconv.Atoi(yearInput)
//
//    if err != nil{
//    fmt.Println("error" , err)
//    }
//
//     fmt.Printf("Year: %d\n", year)
//
//     fmt.Println("Next year:", year + 1)


reader := bufio.NewReader(os.Stdin)

fmt.Print("Enter year: ")

inptstr, err := reader.ReadString('\n')
if err != nil {
    fmt.Println("Input error:", err)
    return
}

year, err := strconv.Atoi(strings.TrimSpace(inptstr))
if err != nil {
    fmt.Println("Invalid year:", err)
    return
}

fmt.Print("Enter value: ")

inValue, err := reader.ReadString('\n')
if err != nil {
    fmt.Println("Invalid value:", err)
    return
}

value, err := strconv.ParseFloat(strings.TrimSpace(inValue), 64)
if err != nil {
    fmt.Println("Invalid value:", err)
    return
}

fmt.Printf("Asset Year: %d\n", year)
fmt.Printf("Estimated Value: %.2f\n", value)

}
