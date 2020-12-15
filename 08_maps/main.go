package main

import "fmt"

func main() {
	// this is saying make a map, the keys a string, and the values a string
	emails := make(map[string]string)
	//this is the fast way
	emails2 := map[string]string{"cricket":"cricket@email.com", "Zoe":"zoey@yahoo.com"}

	//assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Squid"] = "squid@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(emails2)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

}
