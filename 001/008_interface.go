package main

import "fmt"


type Engine interface {
	TurnOn()
}

type Car struct {}

func (c Car) TurnOn() {
	fmt.Printf("Turn ON!!!! \n")
}


type MotorCycle struct {

}

func (m MotorCycle) TurnOn() {
	fmt.Printf("Turn ON!!!! \n")
}


func TurnOnEngine(engine Engine) {
	engine.TurnOn()
}

func useInterfaces() {
	car := Car{}
	TurnOnEngine(car)
	bike := MotorCycle{}
	TurnOnEngine(bike)
}