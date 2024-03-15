package main

import "fmt"

func main() {
	var intSlice []int32 = []int32{1, 2, 3, 4}
	var floatSlice []float32 = []float32{1.3, 5.4, 2.3, 6.0}

	intSum := sumSlice(intSlice)
	floatSum := sumSlice(floatSlice)

	fmt.Println(intSum)
	fmt.Println(floatSum)

}

func sumSlice[T int32 | float32](slice []T) T {
	var sum T
	for i := range slice {
		sum += slice[i]
	}

	return sum

}
