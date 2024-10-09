package main

import (
	"fmt"
	"strconv"
)

// ２つの引数を足す関数
func add(x int, y int) int {
	return x + y
}

// 2つの引数を掛ける関数
func mul(x, y int) int {
	return x * y
}

// 文字を入れ替える関数
func swp(x, y string) (string, string) {
	return y, x
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(strconv.Itoa(i) + "は２の倍数")
		} else if i%3 == 0 {
			fmt.Println(strconv.Itoa(i) + "は３の倍数")
		} else if i%4 == 0 {
			fmt.Println(strconv.Itoa(i) + "は３の倍数")
		}
	}

	for i = 0; i < 10; i++ {
		switch i {
		case 0:
			fmt.Println("aaa")
		case 1:
			fmt.Println("bbb")
		case 2:
			fmt.Println("ccc")
		default:
			fmt.Println("other")
		}
	}

}
