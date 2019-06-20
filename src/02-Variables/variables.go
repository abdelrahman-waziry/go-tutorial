package main

import (
	"fmt"
	"math"
)

func main() {

	var num int
	fmt.Println("First var type is int ", num)

	num = 50
	fmt.Println("Now this var has a value", num)

	var anotherNum int = 58
	fmt.Println("Here's another var with initial value", anotherNum)

	var initialVal = 55
	fmt.Println("No no need to explicitly define variable type if it's declared with inital value", initialVal)

	var width, height int = 60, 60
	fmt.Println("Box width is", width, "and height is", height)

	// Object, is that you ?
	var (
		name string = "Waziry"
		age  int    = 25
	)
	fmt.Println("My name is ", name, "and I'm ", age, "years old")

	// Multiple declaration shorthand
	namo, ago := "Waziry", 25
	fmt.Println("My name is ", namo, "and I'm ", ago, "years old")

	a, b := 145.8, 543.8
	c := math.Min(a, b)
	fmt.Println("minimum value is ", c)

}
