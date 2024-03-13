package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 5
	b := 0
	var res, rem, err = intDivision(a, b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("The result is %v and the reminder is %v", res, rem)

}

func intDivision(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("Cannot Divide by zero!!")
	}

	return a / b, a % b, err

}
