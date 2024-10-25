package main

import "fmt"

type Transmission struct {
}

func (e Transmission) ShiftUp() {
	fmt.Println("shifting up..")
}
func (e Transmission) ShiftDown() {
	fmt.Println("Shifting down...")
}
