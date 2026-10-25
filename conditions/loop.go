package main

import "fmt"


func main(){
// assetType := "equipment"

// switch assetType {
// case "vehicle":
// fmt.Println("Vehicle Selected")
//
// case "equipment":
// fmt.Println("Equipment Selected")
//
// case "property":
// fmt.Println("property Selected")
//
// default:
// fmt.Println("Unknown Selected")


// case "car" , "truck" , "bus":
// fmt.Println("vehicle")
//
// case "exca" , "crane":
// fmt.Println("Heavy")
//
// default:
// fmt.Println("unkown")
// }

// for i:=1; i<=5;i++{
// if i == 3 {
// fmt.Println("enemy ahead")
// continue
// }
// fmt.Println("Asset" , i)}



totalAssets:=5

for i:=1;i<=totalAssets;i++{
fmt.Println("====Processing Asset", i)

switch {
case i%2 == 0:
fmt.Println("Even Asset Number")

default:
fmt.Println("Odd asset number")
}

}

}