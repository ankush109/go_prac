package main

type Startable interface {
	Start()
}

// composition pattern in go
func startEngines(cars ...Startable) {
	for _, car := range cars {
		car.Start()
	}
}
func main() {
	myConvertable := Convertable{
		Engine{},
		EnhancedTransmission{},
		SteeringWheel{},
	}
	myTruck := Truck{
		Engine{},
		Transmission{},
		SteeringWheel{},
	}
	startEngines(myConvertable, myTruck)

}
