package main

import "fmt"

func main() {
	// arrays
//	var fruitArray [2]string
	
	//Assign values
//	fruitArray[0] = "Apple"
//	fruitArray[1] = "Dragonfruit"
	
	// shorter way of doing it:
//	fruitArray2 := [2]string{"orange", "strawberry"}
//	fmt.Println(fruitArray)
//	fmt.Println(fruitArray2)
//	fmt.Println(fruitArray[0])


// SLICES
	fruitSlice := []string{"apple", "tomato", "tacofruit"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:3])
}
