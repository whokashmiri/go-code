package main

import "fmt"

func main(){
purchasePrice := 4511.54100
repairCost := 4554

totalCost :=  purchasePrice + float64(repairCost)

fmt.Printf("Total Cost : %.2f \n" , totalCost)
fmt.Println(10.0 / 3.0)

manufactureYear := 2020
year := 2025

assetAge :=  year - manufactureYear
fmt.Println("Age :" , assetAge)
result :=  manufactureYear >= year
fmt.Println(result)

if assetAge > 6 {
fmt.Println("Asset is Old : " , assetAge)
} else if assetAge == 4 {
fmt.Println("Good deal")
} else if assetAge <=  4  || assetAge == 5{
fmt.Println("perfect")
}else {
fmt.Println("invalid")
}


}