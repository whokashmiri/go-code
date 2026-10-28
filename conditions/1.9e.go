
package main

import "fmt"

type Asset struct {
    Name            string
    ManufactureYear int
}

func (asset Asset) CalculateAge() int {
    return 2026 - asset.ManufactureYear
}

func main() {
    asset := Asset{
        Name:            "CAT 320",
        ManufactureYear: 2018,
    }

    age := asset.CalculateAge()

    fmt.Printf("Asset: %s\n", asset.Name)
    fmt.Printf("Age: %d years\n", age)
}
