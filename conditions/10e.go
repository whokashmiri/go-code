package main

import (
	"errors"
	"fmt"
	"time"
)

type Inspectable interface {
	Inspect() (string, error)
}

var ErrInvalidYear = errors.New("invalid manufacturing year")

type Vehicle struct {
	Name            string
	ManufactureYear int
	IsWorking       bool
}

type Equipment struct {
	Name            string
	ManufactureYear int
	IsWorking       bool
}

func (v Vehicle) Inspect() (string, error) {
	currentYear := time.Now().Year()

	if v.ManufactureYear > currentYear || v.ManufactureYear < 1 {
		return "", fmt.Errorf("vehicle %s: %w", v.Name, ErrInvalidYear)
	}

	age := currentYear - v.ManufactureYear

	if !v.IsWorking || age > 10 {
		return "Detailed Inspection", nil
	}

	return "Standard Inspection", nil
}

func (e Equipment) Inspect() (string, error) {
	currentYear := time.Now().Year()

	if e.ManufactureYear > currentYear || e.ManufactureYear < 1 {
		return "", fmt.Errorf("equipment %s: %w", e.Name, ErrInvalidYear)
	}

	age := currentYear - e.ManufactureYear

	if !e.IsWorking || age > 10 {
		return "Detailed Inspection", nil
	}

	return "Standard Inspection", nil
}

func printInspection(asset Inspectable) {
	result, err := asset.Inspect()

	if err != nil {
		fmt.Println("Inspection Error:", err)
		return
	}

	fmt.Println("Inspection Type:", result)
}

func main() {
	assets := []Inspectable{
		Vehicle{
			Name:            "Truck",
			ManufactureYear: 2020,
			IsWorking:       true,
		},
		Equipment{
			Name:            "Crane",
			ManufactureYear: 2010,
			IsWorking:       true,
		},
		Vehicle{
			Name:            "Bus",
			ManufactureYear: 2000,
			IsWorking:       true,
		},
	}

	for _, asset := range assets {
		printInspection(asset)
	}
}