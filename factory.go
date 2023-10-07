package main

import (
	"fmt"
	"strings"
	"time"
)

type ICar interface {
	start()
	speedUp()
	stop()
}

type SportsCar struct {
	name     string
	maxSpeed int32
}

func (sc *SportsCar) start() {
	fmt.Println("Starting engine...")
	time.Sleep(time.Second * 2)
	fmt.Println(sc.name, "is moving")
}

func (sc *SportsCar) speedUp() {
	fmt.Println(sc.name, "is speeding up to", sc.maxSpeed, "km/h")
	time.Sleep(time.Second)
}

func (sc *SportsCar) stop() {
	fmt.Println(sc.name, "Slowing down")
	time.Sleep(time.Second * 3)
	fmt.Println(sc.name, "has stoppend moving")
}

type SedanCar struct {
	name     string
	maxSpeed int32
}

func (sedan *SedanCar) start() {
	fmt.Println("Starting engine...")
	time.Sleep(time.Second)
	fmt.Println(sedan.name, "is moving")
}

func (sedan *SedanCar) speedUp() {
	fmt.Println(sedan.name, "is speeding up to", sedan.maxSpeed, "km/h")
	time.Sleep(time.Second)
}

func (sedan *SedanCar) stop() {
	fmt.Println(sedan.name, "Slowing down")
	time.Sleep(time.Second)
	fmt.Println(sedan.name, "has stoppend moving")
}

type CarFactory struct{}

func (cr *CarFactory) getCar(carType string, carName string, carMaxSpeed int32) ICar {
	carType = strings.ToLower(carType)
	switch carType {
	case "sedan":
		return &SedanCar{name: carName, maxSpeed: carMaxSpeed}

	case "sports":
		return &SportsCar{name: carName, maxSpeed: carMaxSpeed}
	default:
		return nil
	}
}

func main() {
	factory := CarFactory{}
	solaris := factory.getCar("Sedan", "Hyundai Solaris", 200)
	solaris.start()
	solaris.speedUp()
	solaris.stop()
}
