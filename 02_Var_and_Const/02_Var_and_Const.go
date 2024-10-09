package main

import "fmt"

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
	const Pi = 3.14
	var i int
	var str1, str2 string

	i = add(1, 2)  // 宣言済の変数へ代入
	j := mul(1, 2) // 未宣言の変数へ代入
	str1, str2 = swp("11", "22")
	fmt.Println(Pi)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(str1 + " " + str2)

}
