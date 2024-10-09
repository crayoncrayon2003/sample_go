package main

import "fmt"

func main() {
	//例１
	var arr1 [2]string

	arr1[0] = "Golange"
	arr1[1] = "Java"
	fmt.Println(arr1[0], arr1[1]) //=> Golange Java
	fmt.Println(arr1)             //=> [Golange Java]

	//例２
	var arr2 [2]string = [2]string{"Golang", "Java"}
	fmt.Println(arr2[0], arr2[1]) //=> Golange Java
	fmt.Println(arr2)             //=> [Golange Java]

	//例３
	arr3 := [...]string{"Golang", "Java", "python"}
	fmt.Println(arr3[0], arr3[1]) //=> Golange Java
	fmt.Println(arr3)             //=> [Golange Java]

}
