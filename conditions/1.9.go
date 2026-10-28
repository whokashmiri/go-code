package main

import fmt "fmt"

// type Asset struct{
// Name string
// Brand string
// ManufactureYear int
// EstimatedValue float64
// isWorking bool
// }

type Asset struct{
Name string
Value float64
}

func updateValue(asset *Asset){
asset.Value = 15425
}

func main(){

// assets := []Asset{
// {Name: "Camry" , Value :120000},
// {Name: "Cat" , Value:320000},
// }


asset := Asset{
Name: "Camry" , Value :120000,
// {Name: "Cat" , Value:320000},
}

updateValue(&asset)
fmt.Println(asset.Value)

// fmt.Println(assets)
// for _, asset := range assets{
// fmt.Println(asset)
// fmt.Println(asset.Name  , asset.Value)
// }

// asset :=  Asset{
// Name: "CAT",
// Brand: "Caterpillar",
// ManufactureYear:2020,
// EstimatedValue:20000.500,
// isWorking:true,
// }
// fmt.Println("Asset:", asset)
// fmt.Println(asset.Brand)
// fmt.Println(asset.EstimatedValue)
// asset.isWorking=false
// fmt.Println(asset.isWorking)
// fmt.Println("Asset:", asset)


}