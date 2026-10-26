package main

import "fmt"

func classifyAsset(assetType string) string{
switch assetType {
case "car", "truck":
    fmt.Println("vehicle equipment")
case "bob cat", "excavator":
    fmt.Println("heavy equipment")
default:
    fmt.Println("Unknown Equipment")
}
return assetType
}

func calculateAge(age int) int{

currentAge:=2026
totalAge := currentAge- age
fmt.Println(totalAge)
return totalAge
}
func getInspectionType(age int , isWorking bool) string{
if !isWorking || age >=10{
fmt.Println("Detailed")
return "Detailed"
}else{
fmt.Println("Standard")
return "Standard"
}
}

func printAssetReport(name string, age int, category string, inspection string){
}


func main(){
classifyAsset("excavator")
calculateAge(2020)
getInspectionType(12 , false)
}