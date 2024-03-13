package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	const pi float32 = 3.1415
	var a int32 = 2000
	var b float32 = 5000.877
	fmt.Println((float32(a) + b) / pi)
	c := "ARK"
	fmt.Println("Hello " + c)
}
