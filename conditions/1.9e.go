//
// package main
//
// import "fmt"
//
// type Asset struct {
//     Name            string
//     ManufactureYear int
// }
//
// func (asset Asset) CalculateAge() int {
//     return 2026 - asset.ManufactureYear
// }
//
// func main() {
//     asset := Asset{
//         Name:            "CAT 320",
//         ManufactureYear: 2018,
//     }
//
//     age := asset.CalculateAge()
//
//     fmt.Printf("Asset: %s\n", asset.Name)
//     fmt.Printf("Age: %d years\n", age)
// }


package main

import "fmt"

type Asset struct {
    ID              string
    Name            string
    ManufactureYear int
    EstimatedValue  float64
    IsWorking       bool
}

func (a Asset) CalculateAge() int {
    return 2026 - a.ManufactureYear
}

func (a Asset) InspectionType() string {
    if !a.IsWorking || a.CalculateAge() > 10 {
        return "Detailed"
    }

    return "Standard"
}

func (a *Asset) UpdateValue(newValue float64) {
    if newValue >= 0 {
        a.EstimatedValue = newValue
    }
}

func (a Asset) PrintReport() {
    fmt.Printf("\nAsset: %s\n", a.Name)
    fmt.Printf("Age: %d years\n", a.CalculateAge())
    fmt.Printf("Value: %.2f SAR\n", a.EstimatedValue)
    fmt.Printf("Inspection: %s\n", a.InspectionType())
}

func main() {
    assets := []Asset{
        {
            ID:              "A001",
            Name:            "CAT 320",
            ManufactureYear: 2018,
            EstimatedValue: 185000,
            IsWorking:       true,
        },
        {
            ID:              "A002",
            Name:            "Toyota Camry",
            ManufactureYear: 2014,
            EstimatedValue: 35000,
            IsWorking:       false,
        },
    }

    assets[0].UpdateValue(190000)

    for _, asset := range assets {
        asset.PrintReport()
    }
}
