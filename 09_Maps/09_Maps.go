package main

import "fmt"

func main() {
	// 例１
	map_ex1 := make(map[string]string, 2) //マップの宣言
	fmt.Println(map_ex1)                  //=> map[] //宣言された空のマップ

	map_ex1["Key1"] = "Value1" //マップにkeyとvalueを挿入する。
	map_ex1["Key2"] = "Value2"

	fmt.Println(map_ex1)

	// 例２
	var map_ex2 = map[string]string{"Key1": "Value1", "Key2": "Value2"}
	fmt.Println(map_ex2)

	map_ex2["Key3"] = "Value3" //要素の挿入
	fmt.Println(map_ex2)

	delete(map_ex2, "Key2") //要素の削除
	fmt.Println(map_ex2)

	map_ex2["Key4"] = "Value4"          //要素の挿入
	map_ex2["Key5"] = "Value5"          //要素の挿入
	for index, value := range map_ex2 { //要素でループ
		fmt.Println(index, value)
	}

	for index, _ := range map_ex2 { // indexのみ
		fmt.Println(index)
	}
}
