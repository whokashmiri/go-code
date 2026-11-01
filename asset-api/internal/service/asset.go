package service

import (
	"errors"

	"example.com/asset-api/internal/model"
)

func CalculateAge(asset model.Asset) (int, error) {
	const currentYear = 2026

	if asset.ManufactureYear < 1 ||
		asset.ManufactureYear > currentYear {
		return 0, errors.New(
			"invalid manufacturing year",
		)
	}

	return currentYear - asset.ManufactureYear, nil
}

func InspectionType(asset model.Asset) (string, error) {
	age, err := CalculateAge(asset)
	if err != nil {
		return "", err
	}

	if !asset.IsWorking || age > 10 {
		return "Detailed", nil
	}

	return "Standard", nil
}
