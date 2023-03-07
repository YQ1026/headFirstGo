package main

import "fmt"

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

func TryVehicle(vehicle Vehicle) {
	vehicle.Accelerate()
	vehicle.Steer("left")
	vehicle.Steer("right")
	vehicle.Brake()

	truck, ok := vehicle.(Truck)
	if ok {
		truck.LoadCarGo("cargo")
	} else {
		fmt.Println("类型断言失败。")
	}
}

func main() {
	TryVehicle(Truck("东风卡车1号"))
}
