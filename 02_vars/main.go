package main

import "fmt"

func main(){
  // MAIN TYPES: string, bool, int, 
  // int8 int 16 in32 int64
  //uint uint8 uint 16 uitn32 uitn64 uintptr
  //byte - alias for uint8
  //rune - alias for int32
  //float 32 float 64
  //complex 64 complex128
  
  //var name string = "jordon"
  // but you don't actually need to do that - it infers type from value

  var age int = 30
  var isCool = true
  isCool = false
  
  //const does work

  // this is shorthand
  dog := "Squid"
  size := 1.3

  //multiple assignment:
  name, email := "Jordon", "jordon@jordon.com"


  fmt.Println(age, name, isCool, dog, size, email)
  fmt.Printf("%T\n", isCool)
}
