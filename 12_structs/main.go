package main

import ("fmt"
		"strconv")

// Define person struct
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}
// Greeting method (value receiver)
func (p Person) greet() string{
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) 
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}
// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}


func main() {
	// Init person using struct
	person1 := Person{firstName:"Squid", lastName:"Boy", city: "House", gender: "m", age: 4}
	//alternative (it accepts them as positional arguments)
	person2 := Person{"Squid", "Boy", "House", "m", 4}
	person3 := Person{"Tibby", "ladyington", "House", "f", 4}
	person1.age++
	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	person1.hasBirthday()
	person3.getMarried("Williams")
	fmt.Println(person3.greet())

}
