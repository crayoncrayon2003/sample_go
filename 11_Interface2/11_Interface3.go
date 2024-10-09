package main

import "fmt"

type Person interface {
	Title() string // 敬称
	Name() string  // 名前
}

//type Person struct {
//	*Title
//	*Name
//}

type Name struct {
	FirstName string
	LastName  string
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.FirstName, n.LastName)
}

type Title struct {
	str string
}

func (n *Title) Nickname() string {
	return fmt.Sprintf("%s", n.str)
}

func main() {
	n := &Name{
		FirstName: "Taro",
		LastName:  "Yamada",
	}
	t := &Title{
		str: "tarochan",
	}

	p := new(Person)
	fmt.Println(p.Title())
	fmt.Println(p.Nickname())
}
