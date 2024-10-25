package main

import "fmt"

type EnhancedTransmission struct {
}

func (t EnhancedTransmission) ShiftUp() {
	fmt.Println("quick shifting up....")
}

func (t EnhancedTransmission) ShiftDown() {
	fmt.Println("quick shifting down....")
}
