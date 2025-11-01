package main

import "fmt"

type Address struct {
	Street string
	Zipcode string
	City string
	State string
	Number int

}

type User struct {
	Name string
	Age int
	Active bool
	Address Address
}

func (u User) getFullAddress() {
	fmt.Printf("State: %s \n", u.Address.State)
	fmt.Printf("City: %s \n", u.Address.City)
	fmt.Printf("Street: %s \n", u.Address.Street)
	fmt.Printf("Number: %d \n", u.Address.Number)
	fmt.Printf("Zipcode: %s \n", u.Address.Zipcode)
}

func useStructs() {
	user := User{
		Name: "Zuko",
		Age: 19,
		Active: true,
	}
	
	fmt.Printf("Name: %s, age: %d, online: %t \n", user.Name, user.Age, user.Active)
	user.Active = false
	fmt.Printf("Name: %s, age: %d, online: %t \n \n", user.Name, user.Age, user.Active)

	user.Address.City = "Osasco"
	user.Address.Number = 123
	user.Address.State = "São Paulo"
	user.Address.Street = "street fools"
	user.Address.Zipcode = "020202-020"
	fmt.Printf("***** User Address **** \n")
	user.getFullAddress()

}