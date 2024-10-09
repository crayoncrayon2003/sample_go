package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sum, i int

	sum = 0
	for i = 0; i < 10; i++ {
		sum += i
	}
	sum++

	fmt.Println(sum)

	// while文はない。forでかく
	for i > 0 {
		i--
		fmt.Println(strconv.Itoa(i) + "countdown")
	}
}
