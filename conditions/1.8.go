package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {


// assets:= [3]string{"toyata", "i1o" , "dune"}
//
//  fmt.Println(assets[0])
//     fmt.Println(assets[1])
//     fmt.Println(assets[2])


// assets := []string{"toyata" ,"func" , "dune"}
// fmt.Println(assets)
// assets = append(assets, "i1o")
// fmt.Println(assets)
// assets = append(assets ,"package")
// fmt.Println(assets)
// fmt.Println(len(assets))


// assets := make([]string , 3 ,5)
// assets[0] = "dune"
// assets[1] = "string"
// assets[2] = "ring"
//
// fmt.Println("length : " , len(assets))
// fmt.Println("Capacity : " , cap(assets))


// assets:=[]string{"append" ,"assets" , "package"}

// for _, asset := range assets{
// fmt.Println(asset)
// fmt.Print("\n", i ,asset)
// fmt.Printf( " $%s ",asset)
// }

// for index , asset:= range assets {
//     assets[index] = "Updated"
//
//     fmt.Println(index, asset)
// }



// asset := map[string]string{ "name": "cap" , "brand" : "toyata"}
//
//  fmt.Println("Name:", asset["name"])
//     fmt.Println("Brand:", asset["brand"])
//
//     asset["condition"] = "Good"
//
//     fmt.Println("Condition:" , asset["condition"])
//
//


reader := bufio.NewReader(os.Stdin)

	names := []string{}
	values := make(map[string]float64)

	fmt.Print("How many assets do you have? ")
	input, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil || count < 1 {
		fmt.Println("Please enter a valid positive number.")
		return
	}

	// Collect asset names and values.
	for i := 0; i < count; i++ {
		fmt.Printf("Enter asset %d name: ", i+1)
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		fmt.Printf("Enter value for %s: ", name)
		input, _ := reader.ReadString('\n')
		value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

		if err != nil || value < 0 {
			fmt.Println("Invalid asset value.")
			return
		}

		// Store the name in the slice.
		names = append(names, name)

		// Store the value in the map.
		values[name] = value
	}

	// Calculate and display totals.
	var total float64

	fmt.Println("\n--- Asset Summary ---")

	for _, name := range names {
		value := values[name]
		fmt.Printf("%s: %.2f\n", name, value)
		total += value
	}

	fmt.Printf("Total asset value: %.2f\n", total)

}