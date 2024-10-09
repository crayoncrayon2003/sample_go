package main

import "fmt"

func main() {
	//例１
	slice1 := []string{"Golange", "Java"}
	fmt.Println(slice1)

	//例２
	var slice2 []string
	fmt.Println(slice2)

	var arr2 [2]string = [2]string{"Golang", "Java"}
	fmt.Println(arr2[0], arr2[1]) //=> Golange Java
	fmt.Println(arr2)             //=> [Golange Java]

	//例３
	arr3 := [...]string{"Golang", "Java", "python"}
	slice3 := arr3[0:2]
	fmt.Println(slice3) //=> [Golange Java]

	//要素追加
	slice3 = append(slice3, "aaaaaaa")
	fmt.Println(slice3)

	fmt.Println(len(slice3), cap(slice3))
}
