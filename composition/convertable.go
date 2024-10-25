package main

import "fmt"

type Convertable struct {
	Engine
	Transmission
	SteeringWheel
}

func (t Convertable) ConvertTop() {
	fmt.Println("Converting Top....")
}
