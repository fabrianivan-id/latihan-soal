package main

import (
	"fmt"
	"math"
)

type Vehicle interface {
	SetMaxSpeed(speed int)
	GetMaxSpeed() int
	CalculateFuelConsumption(distance int) float64
	DisplayInfo()
	getLicensePlate() string
}

type baseVehicle struct {
	licensePlate string
	maxSpeed     int
}

func (b *baseVehicle) DisplayInfo() {
	fmt.Printf("License: %s, MaxSpeed: %d\n", b.licensePlate, b.maxSpeed)
}
func (b *baseVehicle) getLicensePlate() string { return b.licensePlate }

// ------- Car -------
type Car struct {
	baseVehicle
	fuelEfficiency int // km per liter
}

func NewCar(licensePlate string, fuelEfficiency, maxSpeed int) *Car {
	return &Car{
		baseVehicle:    baseVehicle{licensePlate: licensePlate, maxSpeed: maxSpeed},
		fuelEfficiency: fuelEfficiency,
	}
}
func (c *Car) SetMaxSpeed(speed int) { c.maxSpeed = speed }
func (c *Car) CalculateFuelConsumption(d int) float64 {
	if c.fuelEfficiency <= 0 {
		return math.Inf(1) // Infinity
	}
	return float64(d) / float64(c.fuelEfficiency)
}

func (c *Car) GetFuelEfficiency() int   { return c.fuelEfficiency }
func (c *Car) SetFuelEfficiency(fe int) { c.fuelEfficiency = fe }

// ------- Truck -------
type Truck struct {
	baseVehicle
	fuelEfficiency int // km per liter
	cargoWeight    int // kg
}

func NewTruck(licensePlate string, fuelEfficiency, cargoWeight, maxSpeed int) *Truck {
	return &Truck{
		baseVehicle:    baseVehicle{licensePlate: licensePlate, maxSpeed: maxSpeed},
		fuelEfficiency: fuelEfficiency,
		cargoWeight:    cargoWeight,
	}
}
func (t *Truck) SetMaxSpeed(speed int) { t.maxSpeed = speed }
func (t *Truck) CalculateFuelConsumption(d int) float64 {
	if t.fuelEfficiency <= 0 {
		return math.Inf(1) // Infinity
	}
	return float64(d)/float64(t.fuelEfficiency) + float64(t.cargoWeight)*0.05
}

func (t *Truck) GetFuelEfficiency() int   { return t.fuelEfficiency }
func (t *Truck) SetFuelEfficiency(fe int) { t.fuelEfficiency = fe }
func (t *Truck) GetCargoWeight() int      { return t.cargoWeight }
func (t *Truck) SetCargoWeight(w int)     { t.cargoWeight = w }

func main() {
	car := NewCar("B-1234-XY", 15, 180)
	truck := NewTruck("B-9999-TR", 6, 2000, 120)

	car.DisplayInfo()
	truck.DisplayInfo()

	fmt.Println("Car fuel for 300km:", car.CalculateFuelConsumption(300))
	fmt.Println("Truck fuel for 300km:", truck.CalculateFuelConsumption(300))
}
