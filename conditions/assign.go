
package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter asset name: ")
    strName, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Input error:", err)
        return
    }
    assetName := strings.TrimSpace(strName)

    fmt.Print("Manufacturing year: ")
    strYear, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Input error:", err)
        return
    }

    year, err := strconv.Atoi(strings.TrimSpace(strYear))
    if err != nil {
        fmt.Println("Invalid manufacturing year:", err)
        return
    }

    fmt.Print("Working condition (true/false): ")
    strWorking, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Input error:", err)
        return
    }

    isWorking, err := strconv.ParseBool(
        strings.TrimSpace(strWorking),
    )
    if err != nil {
        fmt.Println("Invalid working condition:", err)
        return
    }

    fmt.Print("Estimated value: ")
    strValue, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Input error:", err)
        return
    }

    estValue, err := strconv.ParseFloat(
        strings.TrimSpace(strValue), 64,
    )
    if err != nil {
        fmt.Println("Invalid estimated value:", err)
        return
    }

    const currentYear = 2026

    if year < 1 || year > currentYear {
        fmt.Println("Invalid manufacturing year")
        return
    }

    if estValue < 0 {
        fmt.Println("Invalid estimated value")
        return
    }

    age := currentYear - year

    fmt.Println("\n----- ASSET REPORT -----")
    fmt.Printf("Asset Name: %s\n", assetName)
    fmt.Printf("Manufacturing Year: %d\n", year)
    fmt.Printf("Age: %d years\n", age)
    fmt.Printf("Estimated Value: %.2f SAR\n", estValue)
    fmt.Printf("Working: %t\n", isWorking)

    if age <= 5 {
        fmt.Println("Category: Recent Asset")
    } else if age <= 10 {
        fmt.Println("Category: Used Asset")
    } else {
        fmt.Println("Category: Old Asset")
    }

    if !isWorking || age > 10 {
        fmt.Println("Detailed Inspection Required")
    } else {
        fmt.Println("Standard Inspection")
    }
}
