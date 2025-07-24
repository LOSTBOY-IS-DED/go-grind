package main

import "fmt"

type Person struct {
	name    string
	age     int
	contact Contact
}

type Contact struct {
	phone string
	email string
}

func main() {
	fmt.Println("Person struct below : ")
	person := Person{
		name: "John Doe",
		age:  30,
		contact: Contact{
			phone: "1234567890",
			email: "johndoe@gmail.com",
		},
	}

	var person2 = new(Person)
	person2.name = "John Doe"
	person2.age = 30
	person2.contact.phone = "1234567890"
	person2.contact.email = "johndoe@gmail.com"

	// to access the fields of the struct
	fmt.Println(person.name)
	fmt.Println(person.age)
	fmt.Println(person.contact.phone)
	fmt.Println(person.contact.email)
	fmt.Println(person2)
}
