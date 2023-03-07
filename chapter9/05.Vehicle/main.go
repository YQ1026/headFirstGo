package main

import "fmt"

type Car string

func (c Car) Accelerate() {
	fmt.Printf("加速 %v\n", c)
}

func (c Car) Brake() {
	fmt.Printf("停止 %v\n", c)
}

func (c Car) Steer(direction string) {
	fmt.Println("Turning", direction)
}

type Truck string

func (t Truck) Accelerate() {
	fmt.Printf("加速 %v\n", t)
}

func (t Truck) Brake() {
	fmt.Printf("停止 %v\n", t)
}

func (t Truck) Steer(direction string) {
	fmt.Println("Turning", direction)
}

func (t Truck) LoadCarGo(cargo string) {
	fmt.Println("loading", cargo)
}

type Vehicle interface {
	Accelerate()
	Brake()
	Steer(direction string)
}

func main(){
	var vehicle Vehicle = Car("特斯拉")
	vehicle.Accelerate()
	vehicle.Brake()
	vehicle.Steer("左边")

	vehicle = Truck("东风卡车")
	vehicle.Accelerate()
	vehicle.Brake()
	vehicle.Steer("右边")
}