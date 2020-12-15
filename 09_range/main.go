package main

import "fmt"

func main() {
	ids := []int{22,324,54,23,234}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}
	// not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)
	// Range with map
	emails := map[string]string{"do":"do@mail.com", "doom":"doom@gmail.com"}
	for k, v := range emails {
		fmt. Printf("%s: %s\n", k, v)
	}	
	for k := range emails {
		fmt.Println("Name: " + k)
	}
}

