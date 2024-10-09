package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	//例１
	var mike Person

	mike.name = "mike"
	mike.age = 1

	fmt.Println(mike.name, mike.age)

	//例２
	bob := Person{"Bob", 30}
	fmt.Println(bob.name, bob.age)

	//例３
	ken := Person{age: 15, name: "ken"}
	fmt.Println(ken.name, ken.age)

}
