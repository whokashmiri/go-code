package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// var fruitList = []string{"Apple", "ball"}
	// fmt.Println(fruitList)
	// fruitList = append(fruitList, "Cat", "End")
	// fmt.Println(fruitList)
	// fruitList = append(fruitList[1:3], "Fig")
	// fmt.Println(fruitList)

	// highScores := make([]int , 3)
	// highScores[0] = 55
	// highScores[1] = 66
	// highScores[2] = 99
	// fmt.Println(highScores)
	// highScores = append(highScores, 77 , 88)
	// fmt.Println(highScores)
	// sort.Ints(highScores)
	// fmt.Println(highScores)

	fmt.Println("Switch here")

rand.Seed(time.Now().Unix())
dice := rand.Intn(6)+1
fmt.Println(dice)
}
