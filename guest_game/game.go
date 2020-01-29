package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var inputNum int
	var secretNumber int
	var isMatch bool = false

	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(10)

	for isMatch == false {
		println("Please input your number : ")
		fmt.Scan(&inputNum)
		if secretNumber == inputNum {
			isMatch = true
			println("You Won !!")
		} else if inputNum < secretNumber {
			println("too  low")
		} else if inputNum > secretNumber {
			println("too high")
		}
	}
}
