package main

import (
	"fmt"
	"strings"
)

func main() {
	var myStr = ""
	var ark []string = []string{"a", "r", "k"}
	var strBuilder strings.Builder
	for i := range ark {
		strBuilder.WriteString(ark[i])
	}
	myStr = strBuilder.String()
	fmt.Println(myStr)
}
