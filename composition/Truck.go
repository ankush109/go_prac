package main

import "fmt"

type Truck struct {
	Engine
	Transmission
	SteeringWheel
}

func (t Truck) FourwheelDrive() {
	fmt.Println("Toggling 4WD....")
}
