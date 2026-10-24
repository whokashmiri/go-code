
package main

import "fmt"

func main() {
    assetName := "Caterpillar Excavator"
    manufactureYear := 2018
    currentYear := 2026
    estimatedValue := 185000.50
    isWorking := true

    assetAge := currentYear - manufactureYear

    fmt.Println("----- ASSET REPORT -----")
    fmt.Printf("Asset: %s\n", assetName)
    fmt.Printf("Age: %d years\n", assetAge)
    fmt.Printf("Value: %.2f SAR\n", estimatedValue)

    fmt.Println(false || true)

    if isWorking {
        fmt.Println("Condition: Working")
    } else {
        fmt.Println("Condition: Not Working")
    }

    if assetAge > 10 || !isWorking {
        fmt.Println("Inspection: Detailed")
    } else {
        fmt.Println("Inspection: Standard")
    }
}
