package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Asset Name: ")
	assetName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	assetName = strings.TrimSpace(assetName)

	fmt.Print("Enter Brand: ")
	brandName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}
	brandName = strings.TrimSpace(brandName)

	fmt.Print("Enter Manufacture Year: ")
	yearStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	year, err := strconv.Atoi(strings.TrimSpace(yearStr))
	if err != nil {
		fmt.Println("Invalid year input:", err)
		return
	}

	fmt.Print("Enter Value: ")
	valueStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input error:", err)
		return
	}

	value, err := strconv.ParseFloat(strings.TrimSpace(valueStr), 64)
	if err != nil {
		fmt.Println("Invalid value:", err)
		return
	}

	fmt.Println("Asset Name:", assetName)
	fmt.Println("Brand Name:", brandName)
	fmt.Printf("Manufacture Year: %d\n", year)
	fmt.Printf("Value: %.2f\n", value)
}