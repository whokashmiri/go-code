package main

import "fmt"

// var companyName = "valTech"
//
// func printAsset() string{
//     assetName:= "CAT"
//     return assetName
// }

// func changeValue(price int){
// price = 120
// fmt.Println("Inside function : " , price)
// }

func updateAssetName(name *string) {
    *name = "Caterpillar 320D"
    fmt.Println("Inside:", *name)
}

func main() {
    assetName := "CAT Excavator"

    fmt.Println("Before:", assetName)

    updateAssetName(&assetName)

    fmt.Println("After:", assetName)
// price := 100
//
// pricePointer := &price
//
// fmt.Println("Price" , price)
// fmt.Println("pricePointer" , pricePointer)
// fmt.Println("*" , *pricePointer)
// fmt.Println("&" , &price)
//     assetName :=  "Camry"

// fmt.Println(printAsset())
// fmt.Println(companyName)


// age:=9
//
// if age>=8{
// status:="used asset"
// fmt.Println(status)
//
// }

// fmt.Println(status)

// price := 156
//
// changeValue(price)
// price = 160
//
// fmt.Println("Outside function" , price)
}