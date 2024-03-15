package main

import "fmt"

func main() {
	var thing1 [5]int32 = [5]int32{1, 2, 3, 4, 5}
	fmt.Printf("\nThis is the memory location of thing1: %p", &thing1)
	var result [5]int32 = square(&thing1)
	fmt.Println(result)
}

func square(thing2 *[5]int32) [5]int32 {
	fmt.Printf("\nThis is the memory location of thing2: %p", thing2)
	for i, v := range thing2 {
		thing2[i] = v * v

	}
	return *thing2
}
