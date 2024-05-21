package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("The time is ", time.Now())
	add_final()
	fmt.Println(split(3))
}

package elon
import "fmt"

// TODO: define the 'Drive()' method
func Drive(C *Car){
    if(C.battery>= C.batteryDrain){
    C.distance += C.speed
    C.battery -= C.batteryDrain
        }
}

// TODO: define the 'DisplayDistance() string' method
func DisplayDistance(C *Car){
    fmt.Sprintf("Driven %d meters",C.distance)
}

// TODO: define the 'DisplayBattery() string' method
func DisplayBattery(C *Car){
    fmt.Sprintf("Battery at %d%",C.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (C *Car) CanFinish(trackDistance int) bool{
        for trackDistance> C.distance{
            if C.battery>= C.batteryDrain{
            C.distance += C.speed
            C.battery -= C.batteryDrain
        }
    }
    return true
}



