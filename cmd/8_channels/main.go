package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_BEEF_PRICE float32 = 4

func main() {
	var chickenChannel = make(chan string)
	var beefChannel = make(chan string)
	websites := []string{"a.com", "b.com", "c.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkBeefPrices(websites[i], beefChannel)
	}
	sendMessage(chickenChannel, beefChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var price = rand.Float32() * 20
		if price <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
		}
	}
}

func checkBeefPrices(website string, beefChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var price = rand.Float32() * 20
		if price <= MAX_BEEF_PRICE {
			beefChannel <- website

		}
	}
}

func sendMessage(chickenChannel chan string, beefChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Println("Text Sent: found a deal on: ", website)

	case website := <-beefChannel:
		fmt.Println("Email sent: found a deal on: ", website)

	}
}
