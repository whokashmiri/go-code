package main
import  (
"fmt"
"strconv"
)

func main(){
fmt.Print("Hello ")
fmt.Println("Go!")

name := "cart camry"
year := 2025
fmt.Printf("Asset: %s, Year: %d", name, year)

assetName :=  "Crane"
year := 22025
price := 25484.4500
isWorking :=  true

fmt.Printf("Asset : %s\n", assetName)
fmt.Printf("Asset: %s\n", assetName)
fmt.Printf("Year : %d\n", year)
fmt.Printf("Value : %.2f SAR \n" , price)
fmt.Printf("Working : %t\n", isWorking)

result :=  float64( year) + price
fmt.Printf("Result : %.2f\n" , result)
fmt.Print("===========END===========")

yearText := "2023"
 year , err := strconv.Atoi(yearText)

 if err != nil {
 fmt.Println("Invalid year : ", year)
 return
 }

fmt.Printf("Year %d\n" , year)
fmt.Printf("Type : %T\n" , year)

fmt.Println("===========END===========")

value, err := strconv.ParseFloat("12", 65)
fmt.Printf("Value  %.2f\n" , value)
fmt.Printf("Type : %T" , value)
}