package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var count int

	fmt.Print("How many assets do you want to process? ")
	fmt.Scanln(&count)

	for i := 1; i <= count; i++ {
		fmt.Printf("\nEnter asset %d name: ", i)
		assetName, _ := reader.ReadString('\n')
		assetName = strings.TrimSpace(assetName)

		fmt.Printf("Enter asset %d type: ", i)
		assetType, _ := reader.ReadString('\n')
		assetType = strings.TrimSpace(assetType)

		fmt.Printf("Asset: %s\n", assetName)

		switch assetType {
		case "car", "truck":
			fmt.Println("Classification: Vehicle")
		case "excavator", "crane":
			fmt.Println("Classification: Heavy Equipment")
		default:
			fmt.Println("Classification: Other Asset")
		}
	}
}