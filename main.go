package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var tien person

	tien.firstName = "Tien"
	tien.lastName = "Le Manh"

	fmt.Println(tien)
	fmt.Printf("%+v", tien)
}
