package main

import "fmt"

func calTotal(price float64 , repairCost float64 ,)(float64 , bool){
if price < 0 || repairCost < 0{
return 0 , false
}

total :=  price + repairCost
return total, true
}


func main (){
total , valid :=  calTotal(10000,500)

  if !valid {
        fmt.Println("Invalid values")
        return
    }

    fmt.Printf("Total: %.2f SAR\n", total)
}