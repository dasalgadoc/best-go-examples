package main

import "fmt"

// private struct
type car struct {
	brand     string
	reference string
	model     int
}

// public struct
type Owner struct {
	Name      string
	LicenceID string
}

func NewCar(brand, reference string, model int) *car {
	return &car{
		brand:     brand,
		reference: reference,
		model:     model,
	}
}

func main() {
	// Method one
	myCar := car{
		brand:     "Nissan",
		reference: "Kicks",
		model:     2024,
	}
	fmt.Println(myCar)

	// Method two
	var otherCar car
	otherCar.brand = "Renault"
	otherCar.reference = "Sandero"
	fmt.Println(otherCar)

	// Method three
	vWCar := new(car)
	vWCar.brand = "Volkswagen"
	vWCar.reference = "Bora"
	fmt.Println(vWCar)

	// Method four
	sCar := NewCar("Suzuki", "Swift", 2000)
	fmt.Println(sCar)
}
