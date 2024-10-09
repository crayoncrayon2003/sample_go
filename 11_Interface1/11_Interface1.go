package main

import "fmt"

// 空インターフェース定義
var x interface{}

func main() {
	var num int
	var str string

	num = 0
	str = "hello"

	//どんな型でも代入可能
	x = num
	fmt.Println(x)
	x = str
	fmt.Println(x)

}
