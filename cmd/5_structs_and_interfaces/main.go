package main

import "fmt"

type gasEngine struct {
	mpg     uint
	gallons uint
	Owner
}

type electricEngine struct {
	mkwh uint
	kwh  uint
	Owner
}

type engine interface {
	milesLeft() uint
}

type Owner struct {
	name string
}

func (e gasEngine) milesLeft() uint {
	return uint(e.mpg * e.gallons)
}

func (e electricEngine) milesLeft() uint {
	return uint(e.mkwh * e.kwh)
}

func canIMakeIt(e engine, miles uint) bool {
	return e.milesLeft() >= miles
}
func main() {
	var myEngine gasEngine = gasEngine{mpg: 12, gallons: 20, Owner: Owner{"ARK"}}
	var myEngine2 electricEngine = electricEngine{mkwh: 12, kwh: 20, Owner: Owner{"ARK"}}
	myEngine.mpg = 10
	// fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.owner.name)
	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name, myEngine.milesLeft(), canIMakeIt(myEngine, 210))
	fmt.Println(myEngine2.mkwh, myEngine2.kwh, myEngine2.name, myEngine2.milesLeft(), canIMakeIt(myEngine2, 210))

}
