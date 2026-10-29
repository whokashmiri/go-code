// package main
//
// import "fmt"
//
// type Reporter interface {
//  PrintReport()
//  }
//
// type Vehicle struct {
// Name string
// }
//
// type Equipment struct {
// Name string
// }
//
// func (v Vehicle) PrintReport(){
// fmt.Println("Vehicle : " , v.Name)
// }
//
// func (e Equipment) PrintReport(){
// fmt.Println("Equipment : " , e.Name)
// }
//
//
// func main(){
//
// car :=  Vehicle{Name : "camry"}
// machine := Equipment{Name : "CAT"}
//
// var reporter Reporter
//
// reporter =  car
// reporter.PrintReport()
//
// reporter = machine
// reporter.PrintReport()
//
//
// }



package main
 import "fmt"



 type Reporter interface{
 PrintReport()
 }

type Vehicle struct{
Name string
}

type Equipment struct{
Name string
}

func (v Vehicle) PrintReport(){
fmt.Println("Vehicle : " , v.Name)
}

func (e Equipment) PrintReport(){
fmt.Println("Equipment : " , e.Name)
}

func generateReport(r Reporter){
r.PrintReport()
}


func main (){
car :=  Vehicle{Name :  "Toyota"}
excavator :=  Equipment{Name : "CAT"}

generateReport(car)
generateReport(excavator)
}