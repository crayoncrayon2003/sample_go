package main

import (
	"fmt"
)

func main() {
	var data int
	var pdata *int

	data = 10
	pdata = &data

	fmt.Println(data)
	fmt.Println(&data)
	fmt.Println(pdata)
	fmt.Println(*pdata)
}
