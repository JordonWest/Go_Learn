package main

import "fmt"

// Define person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}
	

func main() {
	// Init person using struct
	person1 := Person{firstName:"Squid", lastName:"Boy", city: "House", gender: "f", age: 4}
	//alternative (it accepts them as positional arguments)
	person2 := Person{"Squid", "Boy", "House", "f", 4}
	fmt.Println(person1)
	fmt.Println(person2)
}
