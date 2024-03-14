package main

import (
	"fmt"
)

func main() {

	//! Arrays
	/* var intArr [3]int32
	intArr[0] = 0
	intArr[1] = 1
	intArr[2] = 2 */
	// var intArr [3]int32 = [3]int32{0, 1, 2}
	// intArr := [3]int32{0, 1, 2}
	intArr := [...]int32{0, 1, 2, 3, 4}
	fmt.Println(intArr)

	//! Slices
	var intSlice []int32 = []int32{0, 1, 2, 3}
	intSlice = append(intSlice, 4)
	fmt.Println(intSlice)

	//! Map
	var myMap map[string]int32 = make(map[string]int32)
	myMap["ARK"] = 23
	myMap["SM"] = 22
	println(myMap["SM"])
	var myMap2 = map[string]int32{"Amer": 23, "SJ": 23}
	println(myMap2["SJ"])
	for name, age := range myMap {
		fmt.Printf("Name: %v age: %v\n", name, age)
	}

	/* for i := 0; i < 10; i++ {
		fmt.Println(i)
	} */
	/* for i := range intArr {
		fmt.Println(i)

	} */

}
