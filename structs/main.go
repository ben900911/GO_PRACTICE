package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	alex := person{
		firstname: "Alex",
		lastname:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 94000,
		},
	}
	alex.updateName("Jimmy")
	alex.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)

	fmt.Println(p)
}
