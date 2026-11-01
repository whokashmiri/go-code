package main

import ("fmt"
     "example.com/asset-api/internal/model"
      "example.com/asset-api/internal/service"
    )


func main() {
asset := model.Asset{
    ID: "Asset-001",
    Name: "CAT-321 Ex",
    ManufactureYear:2018,
    EstimatedValue:185000,
    IsWorking: true,
    }

age, err := service.CalculateAge(asset)

if err != nil{

    fmt.Println("Error" , err)
    return
    }

inspection , err := service.InspectionType(asset)

    if err !=nil{
     fmt.Println("Error" , err)
        return
    }


fmt.Println("====Asset Summary======")
fmt.Println("ID:", asset.ID)
    fmt.Println("Name:", asset.Name)
    fmt.Println("Age:", age)
    fmt.Printf("Value: %.2f SAR\n", asset.EstimatedValue)
    fmt.Println("Inspection:", inspection)

}