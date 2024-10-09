package main

import "fmt"

type Name struct {
	FirstName string
	LastName  string
}

func (n *Name) String() string {
	return fmt.Sprintf("%s %s", n.FirstName, n.LastName)
}

type Person struct {
	// *Name型の値を埋め込む
	*Name
	Age int
}

func main() {
	n := &Name{
		FirstName: "Taro",
		LastName:  "Yamada",
	}

	p := &Person{
		// *Name型のnを埋め込む
		Name: n,
		Age:  20,
	}
	fmt.Println(p.String())
}
