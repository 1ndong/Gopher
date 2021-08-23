package main

import "fmt"

const (
	Red int = iota + 1
	Blue
	Green
)

const (
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	SmallRoom
)

func SetLight(rooms, room uint8) uint8 {
	return rooms | room
}

func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room //bit clear
}

func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room
}

func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println(("안방에 불을 켠다"))
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("거실에 불을 켠다")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("화장실에 불을 켠다")
	}
	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("작은방에 불은 켠다")
	}
}

func main() {
	fmt.Println(Blue)

	var rooms uint8 = 0
	rooms = SetLight(rooms, MasterRoom)
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, SmallRoom)

	rooms = ResetLight(rooms, SmallRoom)

	TurnLights(rooms)

	const PI = 3.14              //no type const
	const FloatPI float64 = 3.14 //float64 type const

	var a int = PI * 100

	fmt.Println(a)
}
