package main

import (
	"fmt"
	"strconv"
)

//構造体の定義
type MyCar struct {
	name  string
	speed int
}

//MyCar構造体のメソッドrun
func (u *MyCar) run(speed int) string {
	u.speed = speed
	return strconv.Itoa(speed) + "kmで走ります"
}

//MyCar構造体のメソッドstop
func (u *MyCar) stop() {
	fmt.Println("停止します")
	u.speed = 0
}

func main() {
	car := new(MyCar)

	fmt.Println("speed:", car.speed)
	fmt.Println(car.run(10))
	fmt.Println("speed:", car.speed)
	car.stop()
	fmt.Println("speed:", car.speed)
}
